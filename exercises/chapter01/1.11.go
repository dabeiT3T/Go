
// $ go run fetchall.go https://linkedin.com https://apple.com https://microsoft.com https://wordpress.org https://s.w.org https://fonts.gstatic.com https://gmpg.org https://pinterest.com
// > 0.13s    1561 https://fonts.gstatic.com
// > 0.20s  101187 https://apple.com
// > 0.30s   86534 https://s.w.org
// > 0.81s  118671 https://wordpress.org
// > 1.30s  156376 https://microsoft.com
// > Get "https://linkedin.com": net/http: TLS handshake timeout
// > Get "https://gmpg.org": dial tcp 104.244.43.57:443: i/o timeout
// > Get "https://pinterest.com": dial tcp 98.159.108.61:443: i/o timeout
// > 30.00s elapsed
