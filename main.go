package main

import (
	"net/http"
	"fmt"
	"strings"
)

func main(){

	client := &http.Client{}
	resp, _ := http.NewRequest("POST", "https://www.netflix.com/es/login", strings.NewReader("email%3Ddavidestebanell01%40gmail.com%26password%3DEstebanell01%26rememberMe%3Dtrue%26flow%3DwebsiteSignUp%26mode%3Dlogin%26action%3DloginAction%26withFields%3Demail%252Cpassword%252CrememberMe%252CnextPage%26authURL%3D1488932591632.VQJ%252F7fYpQ999sQ4WuB%252FAp7hrGDA%253D%26nextPage%3D"))
	resp.Header.Add("Host", "www.netflix.com")
	resp.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:51.0) Gecko/20100101 Firefox/51.0")
	resp.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	resp.Header.Add("Accept-Language", "es-ES,es;q=0.8,en-US;q=0.5,en;q=0.3")
	resp.Header.Add("Accept-Encoding", "gzip, deflate, br")
	resp.Header.Add("Referer", "https://www.netflix.com/es/login")
	resp.Header.Add("Cookie", "memclid=17a3008b-f36c-40fc-901b-e404f0608b48; nfvdid=BQFmAAEBEDib06ZRoKogEJOJfboUQPxg7pgIUgMSc1HOb5tfX5FC1ZfixpD0Z0wioF3LaALK5KnXOVg1vBX%2FeP93yQAHfLbKX4iWgFAVPmjf8Xh8ON6LGMHUdb8itUDa4OFF2zlVWfRUG0Blj92wYRoT20ufWLbL; dsca=customer; NetflixId=v%3D2%26ct%3DBQAOAAEBEGoQ4CdSErQDmeuHunASCyyBsGlJlg2mQpVc_UiVki_CsJHhykvbd0tFC3BHkNeFh8HfrXLMSNNI7ShZPQ64GNdlaA_N50hGzDp1g3neDTktWDz27-eR1sYKzEoy1OAa9382N6p2jF7iVHXBSZwjYRHcm-TMsc-sjG5hNArhJRbJSqYXv5N5tDlBGCknfUUUGBERSgXVbQ970uLV26JyTcu_rGoqh0D_twOUFceYA5ey3UtlKYSVnOtmmC3PgxgDqQbtChRUsKjz2EZzlzLQFg9_r9Q4DGhdvA4eSQgVuENEhIa9bCwebHUaJ17yWoQ-z47ozPPnLQ95OyetFQKuTLothXAGJvfCIlmi-i9HQcSz_-I_jDFXgo16TjRZwWtMw-hiV9dQc988RHV-aCCxLHksCtpAaSPDECVWelvZwLpIMwgBnFAB4mj3JLsR3Ga4jKw1cUPpVFuUOs3WX2VcLvZsMzlurZvGJvyIhMgRHwBggegIoubJBCMuZMv-zADxVWunc_s-qzUJ1Ks6YwPcYVfOwx2HbNo_-hTRHXdIyWkRCQgSQqnsVj1gHyCLrw3CtyObRyJJlDSoeZNksVPLuH-6PQ..%26bt%3Ddbl%26ch%3DAQEAEAABABRXCis0xTVYlms4-lRsEwUsErTm2X3NW8U.%26mac%3DAQEAEAABABQjNTFEDvgJS7r161S7Ds0QUpAJw6Ds7zU.; SecureNetflixId=v%3D2%26mac%3DAQEAEQABABSGZChDGWh4joK4-C0yFMgxz6clIl6jO0A.%26dt%3D1488931836064; lhpuuidh-browse-OGJ3G6DUKVB6JANSLOXXTXFK2M=ES%3AES-ES%3Ab34cce70-a55c-4f19-a96f-0a2a6f1acb00_ROOT; lhpuuidh-browse-OGJ3G6DUKVB6JANSLOXXTXFK2M-T=1488930852434; profilesNewSession=0; cL=1488931843478%7C14889318367854785%7C148893183691807737%7C%7C9%7CFQI374GDIZD33EE4JB3QKC6ZXY; lhpuuidh-browse-WSHJHEIYAFBZ5K3GB44CFYM43A=ES%3AES-ES%3Adb61ef05-a06d-43cf-bfe0-26fda4265de3_ROOT; lhpuuidh-browse-WSHJHEIYAFBZ5K3GB44CFYM43A-T=1488931115072")
	resp.Header.Add("DNT", "1")
	resp.Header.Add("Connection", "keep-alive")
	resp.Header.Add("Upgrade-Insecure-Requests", "1")
	resp.Header.Add("Cache-Control", "max-age=0")
	response, _ := client.Do(resp)
	fmt.Println(response)
}

