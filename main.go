package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func handleTunneling(w http.ResponseWriter, r *http.Request) {
	dstConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	cliConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(dstConn, cliConn)
	go transfer(cliConn, dstConn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

// setAssumedContentType assumes the Content-Type from the Body given.
// XXX - not implemented yet.
func setAssumedContentType(body []byte) (string, error) {
	return "text/plain", nil
}

func handleHTTP(w http.ResponseWriter, req *http.Request, contType string) {
	// Content-Type handled here.
	if req.Method == http.MethodPost {
		currType := req.Header.Get("Content-Type")
		if currType == "" {
			/*
				l, err := strconv.Atoi(req.Header.Get("Content-Length"))
				if err != nil {
					log.Fatal(err)
				}

				buf := make([]byte, l)
				n, err := req.Body.Read(buf)
				if err != nil {
					log.Printf("Failed to read Request Body: %v", err)
				}

				newType, err := setAssumedContentType(buf[:n])
				if err != nil {
					log.Fatal(err)
				}
			*/
			req.Header.Set("Content-Type", contType)
			log.Printf("Set new Content-Type: %s. Src: %s, Dst: %v", contType, req.RemoteAddr, req.Host)
		}
	}

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func main() {
	var (
		addr     = flag.String("addr", "0.0.0.0:55555", "address to serve HTTP")
		contType = flag.String("type", "text/plain", "Content-Type to set")
	)
	flag.Parse()

	server := &http.Server{
		Addr: *addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				handleTunneling(w, r)
			} else {
				handleHTTP(w, r, *contType)
			}
		}),
	}

	log.Printf("Started listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
