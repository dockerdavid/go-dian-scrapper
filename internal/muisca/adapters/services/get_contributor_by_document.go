package services

import (
	"crypto/tls"
	"errors"
	"github.com/PuerkitoBio/goquery"
	muiscaDomain "github.com/dockerdavid/go-dian-scrapper/internal/muisca/domain"
	"io"
	"net/http"
	"strings"
)

const url = "https://muisca.dian.gov.co/WebRutMuisca/DefConsultaEstadoRUT.faces"

var (
	ErrNoResults = errors.New("no se encontraron resultados")
)

var transport = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

var client = &http.Client{
	Transport: transport,
}

func (s Service) GetContributorByDocument(document string) (*muiscaDomain.Result, error) {
	var data = strings.NewReader("vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AmodoPresentacionSeleccionBO=pantalla&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AsiguienteURL=&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AmodoPresentacionFormBO=pantalla&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AmodoOperacionFormBO=&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AmantenerCriterios=&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AnumNit=" + document + "&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AbtnBuscar.x=58&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AbtnBuscar.y=13&com.sun.faces.VIEW=H4sIAAAAAAAAAO1cW2wkV5mubtsztjOZJAwMhMSkJpOMEzFut9tz8T3x2DPYu74MHs8sEEHndNfpdg3VdSqnTrXbk4sCD%2BzDPrDSggQosLvSvoDgZVnYlUBCaKRdaaWsRNC%2BrLS7iBeEgHB54fIC%2F3%2Fq0lXd1W1XT5lkoqmHcnVVnf%2Bc%2F%2Fuv569z%2FPVfKAM2V2YYr%2BbITdI4d9Ou5IhlGXqZCJ2ZuUW4eU0QQdeJSaqUP7HNKb0muCMcTjeYRj%2F92V9%2B4Xbl9NAxRWlYtxSlX1FOR6iVWc1iJjWFpHVDp7tbjAllqFj3L90jo4zeJHXSyFVImdqhZtdXN0hNN6tLzBRENylXTtZ1WxD4bTuGIJfhWmNb17c9QgPKpTLLVVk9p%2BnEzNUc3S6THOEvOLqgZRg3ye3SUi4yNuhweUXUjKvEpMaHuK4p%2FUVdyyv%2BcblnipuOsBxxhfEaEZJo4TCITgZE53smiuSU91Tg3AnZQWW5Z%2BqrJox4Rdc0aiqP1JjGrnJqw3NSBjW7Rg1axotLmwEj6XR1zNarjo4Pr2%2BtpUz7ZCsbiGDqHJzAXjYtyg%2Bvi4dqBB6AZS1xaMh1ZgcdLCbvQDoBr5d1atrkJrWlkp4LqKZloOc9ev2pUbwQeKOlO7TObdpwbfNi6mxPpc72dEAxFbYHgOREPgW%2BpZZKkkdMp7ahi4BmeuOcSJdmVqunKnA5Rj9oDKfI92S6fN%2FHyS3w5KysE0NR3EwgbVGdS3fIxy2u1yhftKhh6Bo7BN06ny7NB2xadUyNHeKIL6RL85gL8QarlTg9hOFeTHm4THBmu6NtBsHUzLjpuNNDYDp9VAv5dGhKJxAhnLKvPUJlhpq6oAq%2Bv01RUIWU%2Fa2kmbJDlDRT9lmSZtOrpCaipuWnNI%2BSVKfSH6hvoH3KlZ5JLrEa5OjaJUcIZipDJWFecqA1D0abFmkc8WTT%2FHufX3hk1%2FRdw%2FykMgynNb1m6aEhpyi2ycCvNByunHhuDSsJOYOY1dxm6SaQm%2F3b%2F%2FrI3z9oP21ksUwB72WdF5RXlGyHq%2F7gatDmykOSmiN0I7dC7J11Yg0c%2FZ%2Fv3z75%2FA%2F6lOwV4I0R7QopC8ZXlSGxA5Fjhxlaw3rmWTmkzO4gEscroTwRrnEAg8SyckvXt7Yub2wXb6xe%2Fqvi1ubmNnY%2B3LDCRZLwdTAkyeAlxgxKzNdV%2Fup%2Fv%2FaHN7NK5mPKQJ0YDm1YGRtJ3a9YQGt4ZXt9rXhp8drqklAeGV%2BmlbZpfg5YpDaQf7DJ8RorE4O%2B8ruHnn8t%2F%2Fuf9ylHVpXBHUChzDS6phwtM8cUfE8o75KYj%2BOQxq8JrpvV2TVlGBwAzABh3ioHcmJNGcQXHFKl3u8jdpnrlvB%2BHa0TDtrg%2FmxYf4QDEuqlTUj74S%2BVRE7KUyOQ0ABePCBPVvMyeJyJEe%2BgJDQYJx81Tj6b61c3N1BCq8vXAJ0TTXQWOSd7a7otGp96Y%2BSL%2F0G%2B3KdkVpV%2BW79FpZoN7GIWPOBOBPNeAcgr2YgOxRZfpRuoAR1KXR0fIGPHXTCsJkQxGDy2Pxp4e0Q%2B%2FIB80O9ChIN3B%2FdI%2FBhmkFdspjZHI1yd8pFF72hZUUVeBQOvUv6uH%2F%2FDP%2F32U389lUUkPUX2NVK%2Bt%2BHUSpR%2F5uufH7nvcz%2F6G2nQr8LIPKYtOeC8P%2F7jVuzRVA8rBpsjwVVfB5QUidKD0uN0B6KAzU6FgTiFPo1TU6Mc%2FFnYl23Jm5SDufZ9dk6QkkFV2wE%2FyvfmR0fVskFse37UIppGtTH5eFTd1TWxMz86kc8%2FOaraYs%2Bg86MlxoHOTH52dGF4aGhI9Q%2B4nhMlpu213Yb7vP0m3NX8Xmt6ldhFi5hBLxVmirEKqenG3swimK1xVl2hRp0KvUzOqjYx7TGbcr0yq8o30SZmJgpWI2ZU2JUNtKOdaQzgGF04o7%2FgQJ43u3yGkbIj6CzcVeEG23tmRp0bx3YHJKibellnowtzRAUvXZkfdW%2FkdiCmpMHWqiQ3N04WEg3LphbhoDF8dOEln6G40QtaG124bjsEK3Qqp1VQOmxnd%2BxtblxoMbIejxP2%2FhogQ3SRcEqKFgWZmCRekjsTC54poPqqbo6uYnwB7uBhbJvCwpbLEFNvfwlwZOo210uOQGZz0KxwaOx5ckezgqg1ZtCKmJnKWw31FCQsjAuIRyDbILQcD7sP5ZC9xyQ2e7zpPfA0a1lBHDzYiE4m9fWD0tcPiq6VehEtrotO9XARW8IWMVVnN7acc%2F%2Bc9zLyaRlqznbAKDaCIhunwy734YjLxQE0XS3mRqdCn7zGG2O7u7tjSHjM4QY1MdPROkeQwa6ZZDP9GE6kFiFSLQmtl1x9440bP%2FnpyIsf8hNaCMxusGx2DeH1TPwnNflFbwWSVMqvkTrlH%2F3Pb87%2F3Wuvr2eV7JoyJO19g9T8HG3Yhnc02Saa53m5NcrnFBDP2Y7pdUWN3A0czCXdRKNarVmGUMbKbQlnV%2FVaSSLzmS6kkI%2B%2FCOvDIxF9cL%2BChDSimUY2bSyUZXo5xv1KV6NLVQfw6hO%2BnT4fEjFelfAEovlAO75RA51OBGi4Lfbwl1Ecdu5eoJ7aXxF9J7V0RzroUsFe194x4D0ZD16be382MXItJLC%2F9XcMbKdjYGuPfwvJQGslgH1tHBJkzSn10YOlNROJeMHAj8Q2w356JOKn%2FY%2FJgacOXo09uXEakgkDxhFBYZ90KfHUuF%2BmS%2F1u3nLB%2FTPlpS8X%2FcJioydMzmMfV6MSGw9NcuW097hQMtBtJu%2Fdf1VpTodbDki7sC%2FHwLy66Kb0eLfvQj4fU%2F5Kq15wsTfuLyCxb%2B3HPZ6%2Bg6fvHoT%2Fd7fxXxSaqx3f61o1uoMMPzHjF5HYv4dNIZrCYmU9VC1Qnr7KuFohdTjbTg2mtTCRoqoB08QgituqRoTrInyax8rMwJnjTKFhzR7MC9yZbUx736wK3l%2FfSAq9qccUdvLtDuoRVoIDqMV9oeltOL1PxRYCENyFEcFqBn8Fglw10BsG09jH%2F%2B6PgVA%2BWGGmxrZ14RhsDSa5izcdSeis0C12RTfIBhahjBuUazSwsP6QK%2B04BYrDqFdrSRpXT3JHFGteZIAZUokaRRORLSRFckLWLH8WhfKHXuDwa6oHMpKjMZpyuNnF%2B9qzC1%2FDJhPh4LZCqj%2BPB6JrHhGNsoESheOxfzTfC157wItlw7KmOTNxPj8rTXOIY7WJ6xrDQAXD26%2FekaLuZYUy6HpdqkUFkI0I4NF2AWh19dS8ajqGEXnzPTEKq9V70dcJpPiLFPQ1TbzgIjOSDlLvjntTKOOJgNLqSO3NbiiBhrkaN50%2FrCg4LAPAsL%2FyKrpayl%2Fc1LYgyV9E1L7wx1%2Bs07LAxl8U07KQRUaW5PolP1383%2F6hBf%2F8%2BiARA598L8YjvkWa%2BOF2%2FQpJxVdJ9cwZNca7SmSXWM0yqGDqfIz6PhZj6JJ%2B0fbEnlwkWA%2FOzN1tJn%2FIQI90JS%2BUqUQwh5oiY%2FMHAjvztga7BcOueIcFEwf24zFa7bqgIgmcU3LFxrl%2FZuEdoNhpYh1TM4qSfy7%2FcaHMJgI7GmGQw2fuaXgU9RgN98LvHak4lnIyz95T8YQqPgEqPpfsw0U0V0IWF%2B%2FpeIS82tmLm14mmVzBsVqXuXRPwRMqeOHjSb%2FMhVN%2B5G%2Fpnnbvq91yPuQpt92LdmNJNrN8T7sTavdkYu0Oz1yRv8spaPcdVmyn%2FX0R%2FjaGYOtBj1NrWbn%2B%2Fyhj%2B0%2Bts7LBAabWh1yQTbPIMxJjqxVaLpKycHqbI2NFPPPLu81QEyDW%2FiYAtkMWe0WsgIXvzK9aEXM7FcqT0TVEZp1ykVsmgm5DFFpyf1PexEooRzVaIdANmMYJTRtbXx%2Fbg0NdWZmp1WZspJt9H36Y1YDI21wsMasdBG0I5v6Mlc%2FDMRrtOoteZINF3syv7zZt7gG2GLUONUn6HcMFHMf5m7c8fBQm%2Fd1a%2Fg4rf1dUb9GjIAuzP0oaPQ60Hv7P8CkvMxI8iHzEynysXQN0jXo%2F%2FMXBkIVMFCYn8pOT0%2BpLL8WkLF2anMtDk2E15hgaSkZoInnfhbT6Ppe4b1TA5FqGteZs%2B3ROKCToAUKOXmcz6hLjkKxZuF5fJXLZwcbqtvqCQ1WbqtQsO9QUnJhqXa%2B6SxGoqRpELRF4DE3kwgS8gHvLq4sbubfSuXVTUWU8Ee7TF3rBHUuh2fYpnciYQQ%2FXHNuipqZrHbEvO3tE1U25u6qs3%2F6mGZGEKigubSdGDcWh2j45claF2ypuHzHVm46my8SfqUTz1pQQkDcBSZUN3LFA5NvEEYxDW03FBcZUSJIMRkBMcPU1ivv74A7goAtsq9Zx%2B4ReJsG4EFRcXSw1BnRA0zku3cWnuolA%2BV1Rw%2BMJ5i0mU2kDxoR94bUNb9g6bplxSkCcGwy1DFporAZ8GDoLiLGc%2BtQiJDHqxHkYj6Eu0zLHqVDh3IU8DrCQn5h8%2Bp2ihBh1kishFiuz7bNdoVSCHtQlYpapQTorYbsD0E0iXcbd6wFKyYLUVGJnPX2%2BF3lh7S3bPlMWGYaC2GAqp4HlAPTrcrFTi9wMKTTNXTdWNrxVYyhEE80ZJvQ65cxEdlBOrGToVc%2BIoZFLvqRzKXE0PBgbPDCRO4%2BoZZBbKHtbbl4rg%2F%2ByvXcDE7w4NYVvkyqzBXONMX9O3ZPmDyaNu3GkvUf1yp0t6rcAETunqoCNp4LwpsU0fvsfTUBAvsFVUzerjkncRu74pYcCnpqOzr0v%2FD1IRN0hMGqJhoSCuDggK2Ekkitup9yyZQlhcpWQawj%2FtUOWqHhZ4nErdKdThog7eNo3f7UWHVLZC5WcTVlK%2Bbcom7M45oG58RJPsG8rcZLfJ5P8PhHaj%2B%2Ft3c2L8I73HrN8ubLvw%2Ftn%2BU35daoPwQjByEFBgXoYhb4YFEZiZNdlv%2FQxt9y3KHfaCOX9MPfNLS%2FVNLiBXPnubcsRsZ4W8NJrpEo7ONvLOJWWL5gwlXa4USQlmxlgkcHdl0%2B%2FGHnJvSiWmGBw9oRyIRH4gTTRp77WFEDrviH3XyYEy3D9esnjbVuh1qnYYVpoL1RzugWKgtHB7TzC%2B0NtUQPk%2BNjpF9uDSZPAy6FldUI54nJhRZe%2ByzXQ4GsdDvk4eESN8A1dPPX0bEs1Mfz%2BEdsp1XTR1XwSKQ4i%2B1U8fa1TfSH7jWgUjtYXrrQpBodMBJPTgymG4VrmwStwPRjwJJbVsl%2BOGHD2n9vQbcX6vd7YQCKelCrEsGln8SDZ73YVzaO9iEY5iW7rYiKumz4PqXwlbDpqxHSC%2Fw5ifrLFfrDhv4Tt42icrFrn0HG28sHTL%2B73Vs6j%2FXILpDEH7iOhFREplqQS8R7tsvtXrvB8Iia2XYnd%2Bhy38xnute34x5uYgHUKjtafAO%2FaO64XVgAA&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT=vistaConsultaEstadoRUT%3AformConsultaEstadoRUT&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3A_idcl=")

	req, err := http.NewRequest(http.MethodPost, url, data)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Host", "muisca.dian.gov.co")
	req.Header.Set("Content-Length", "6458")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Sec-Ch-Ua", `"Not/A)Brand";v="8", "Chromium";v="126"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Accept-Language", "es-419")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://muisca.dian.gov.co")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.57 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://muisca.dian.gov.co/WebRutMuisca/DefConsultaEstadoRUT.faces")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "JSESSIONID=C0760ECC9C4E6BB9E845C6E251CB4650.nodo11Rutmuisca; DIAN-MUISCA=N_1_393939_19052b085c1_N_61736446313233_; TS01a45c3c=01615e3645a137ff080766f385de2d7bdd4663cdbece5d2885f6e4b130cb880ccbb34992709cfbb2b9c1c04a1b9a863bb3765932bd; TS01a45c3c031=0184d7998b9279e332ca5082f9e84de88bbbc62ab6d03ec2b66ecbf68d04b120eb38670d9716ced785b260bfb4ae9a5af615366ee9; TS01a45c3c028=0184d7998b6edf095d71090dcfdda147d64d9efc5fd03ec2b66ecbf68d04b120eb38670d9761a0b976d7fd7b7aa279b8374066812d")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(bodyText)))

	if err != nil {
		return nil, err
	}

	var tipo string
	var razonSocial, estado, primerApellido, segundoApellido, primerNombre, otrosNombres string

	var table = doc.Find("table.muisca_area")

	if table.Length() == 0 {
		return nil, ErrNoResults
	}

	if table.Find("span#vistaConsultaEstadoRUT\\:formConsultaEstadoRUT\\:razonSocial").Length() > 0 {
		tipo = muiscaDomain.JuridicalPersonType
		razonSocial = table.Find("#vistaConsultaEstadoRUT\\:formConsultaEstadoRUT\\:razonSocial").Text()

		if razonSocial == "" {
			return nil, ErrNoResults
		}
	} else {
		tipo = muiscaDomain.NaturalPersonType
		primerApellido = table.Find("#vistaConsultaEstadoRUT\\:formConsultaEstadoRUT\\:primerApellido").Text()
		segundoApellido = table.Find("#vistaConsultaEstadoRUT\\:formConsultaEstadoRUT\\:segundoApellido").Text()
		primerNombre = table.Find("#vistaConsultaEstadoRUT\\:formConsultaEstadoRUT\\:primerNombre").Text()
		otrosNombres = table.Find("#vistaConsultaEstadoRUT\\:formConsultaEstadoRUT\\:otrosNombres").Text()

		if primerApellido == "" && segundoApellido == "" && primerNombre == "" && otrosNombres == "" {
			return nil, ErrNoResults
		}
	}

	estado = table.Find("#vistaConsultaEstadoRUT\\:formConsultaEstadoRUT\\:estado").Text()

	var result = muiscaDomain.Result{
		NIT:             document,
		State:           estado,
		ContributorType: tipo,
		NaturalPerson: muiscaDomain.NaturalPerson{
			FirstName:      primerNombre,
			MiddleName:     otrosNombres,
			LastName:       primerApellido,
			SecondLastName: segundoApellido,
		},
		JuridicalPerson: muiscaDomain.JuridicalPerson{
			SocialReason: razonSocial,
		},
	}

	return &result, nil
}
