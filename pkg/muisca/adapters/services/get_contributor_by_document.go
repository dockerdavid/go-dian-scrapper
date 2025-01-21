package muiscaServices

import (
	"crypto/tls"
	"errors"
	"github.com/PuerkitoBio/goquery"
	muiscaDomain "github.com/dockerdavid/go-dian-scrapper/pkg/muisca/domain"
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
	var data = strings.NewReader("vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AnumNit=" + document + "&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AbtnBuscar.x=24&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3AbtnBuscar.y=10&com.sun.faces.VIEW=H4sIAAAAAAAAAO1cW2wcR3btGT4kUrJsrXa1a8uyS5IlWog4HA4f4ttLkaLJBR9akpJ37Tjjmu6aYcs9Xa3q6uFQthe7%2B5EgyMcGSAIkgYPEQD7XeTnJf6CPAAE2QBbIT76STYBNkBeSnzx%2BNnWrH9M90zNkj5pZWdB8zPT049a9596699atqv7%2Bvyp9NlNmKKvk8H1cH79vl3PYsgxdxVynZm5RnNzhmJMNbOIKYa%2FtMkJ2OHO4w8gm1ch3v%2Ffvv%2F6ofGXgtKLUrYeK0qsoVyLUVFq1qElMLmnd08n%2BNqVcGSjW%2FEP3k1GG7uMarufKWCV26LG7a5u4qpuVJWpyrJuEKedrus2x%2BG87Bse3xbFGt%2B%2FueoT6lVsqzVVoLafp2MxVHd1WcQ6zB47OiSr4xrl9UspFeBMNLq%2FyqnEHm8R4k%2Bma0lvUtbzif9KiWAgo3u6a4pbDLYevUFbFXBIdPw6iEwHR%2Ba6JAjnlS2Xx3U5Xp5TlrqmvmYLjVV3TiKlcqFKN3mHEFtexKgx3hxhEhYNbW4Eg6TR12tYrjg4X726vp0z7fLMYgGDqEpyDVrYswo6nCekEoJ2zVSyuig67xMTTTKd2yoKc3NO0Xfq%2BOEqX7vOCrmHgGrnjlMATBuQXuya%2FQUwb3ye27FuTqfqVQehgjoEFwB7V3tT81c3AOS89pmvZJXXXsUyl7lSnPXp9qVDsExRH8wGSqcgtaY6mILi0Ukmy33SqmzoPaKbHZyENmtIJNAhntVqqapeMjnkEs%2BmR9ENpT4p4ThyDjiaPgebNY6CZbmeXJP3ePpgem4V8uqKfYvihSEGoqmMjfVALo%2BnSPGMxvUrYokUMQ9doisSlE%2FC5TsWtNLh%2B3iYVx9ToMbDtczyWLs3TLs6btFpi5BjYHU%2BZXcoZtV1uG4lbWt244HvFFCNs4Ri8YuEYvGJhKl2a%2FUSOqtLXke9qH0tHTZkAEB5L2d9Kmil7RUkzZZ8laTa8Slp6Gku9ACCpTqTK6AkVW1zdw%2BkSLVEuLvnuqVdZ6ZroEq2K4ap2y%2BGCojJQ4uYtRzzNAn7TIi2xbXiV7kesHtl13XxfGTTM99f1qqWHGE6H8oCg3IREF0a2f7mtmQXOsO4w5dw761AEzBnYrOS2SvcFwdlf%2Fstv%2FPYL9nUjCxVGcV%2FWeaB8SyT98Ue9wdFJmylnJTWH60ZuFdt7G9jqO%2FE3f%2Fbo%2FHt%2F1aNkVwRsFGsrWOWUrSkDfE9Euj1qaHXrja9KljL7J4E4HHHltXB5UoiILSu3dHd7%2B%2FbmbvHe2u23ittbW7vQ%2BGDdCtc3w8cBS1LAW5QaBJs%2FQOzbf%2F3x%2F%2F5bVsm8rfTVsOGQupWxgdRziiVoDa7ubqwXby3urC1x5cLIMim31NNyQkRiC%2FIvNCRepyo2yLf%2B%2B%2Bx7H%2Bf%2F5196lP415eSeQEGlGlkXHZI6JmcHXPmCxHwEWBrZ4Uw3K7PryqBwWMS0ddEQMHJuXTkJNzi4Qrz%2F%2FbbKdIt7%2F07UMBP24P6tWz8RHzHsu7sjRqjil5hw%2Frz8qgca6oOD5%2BWX1TgMLmdi1HtScnMyTj8oTj9bG3e2NkFDa8s7Ap1zDXQWGcMH67rN69%2F54cXf%2BHP8Wz1KZk3ptfWHRJpZ%2F34vfLtFi7z7U%2FAKrl6JlLcpbvqWXQdDaFOsbnsB5DvjYmIFYvfHQPHq4aDA6Yvy4ivyQq%2BLFDA%2FKpm7EM%2FDDIgMj6EGN9w1LR9g8L%2BWFbXnNdHTK4R94Ue%2F87v%2F9Z1fmMoCoJ49%2B4Yp79t0qiXCfv77v3bx1K%2F%2B7S%2FJfv1twZkntCUZHg2fifk0rMRKHZuxw7ApwGOXGtjAVz6OcfeM39yZbkRpqL6njVCKFOoF6UM78z0Oj10O6%2FQS%2BGlGTI0w4aHD3nlbniRMOKCe781xXDIIsh0RGNjB%2FNAQUg1s2%2FNDFtY0og3Ly0NoX9f43vzQaD5%2FdQjZ%2FMAg80MlygSdmfzs0MLgwMAA8j%2FieI6XqHbQclqcZ60nxVnNb7WqV7BdtLAZtFKmJh8u46puHMwsCkdk3ECrxKgRrqv4BrKxaQ%2FbhOnlWSTvhF4%2BM1qw6jFcQVO2oB1tTKMCjqGFa%2FoDR6Tbs8vXKFYdTmbFWSRO0IM3ZtDcCDx3RIK6qas6HVqYw0jEnfL8kHsityeiZBpirUlycyN4IRFbNrEwExbDhhY%2B9AWK456T6tDCXduBsrONGKkIo4Pn7LatzY1wLUbXI3HKPtwCZNpRxIzgokWETkwcr8m90QWvK4D5IneohCBiCunExdhnCgvbrkAUPfpNgSNFu0wvORyEzYnHCscmnqd36FYiDg8bpMxnpvJWHV0S2R1lXERYodvA6ZwJuw%2FlmL3HBDx2Jer1ViwriOxH4%2Bh8Utd8SrrmU7zjJB%2BPzsvxdlNpPHb2i8fMVfHG9BJvmRFy48QkD0%2B%2B8GCMw4OBiQwkN9pAGps7gNSvhT30ixEPDfw2PDMkh5dC0%2FUj9eH9%2Ff1hIDzsMIOYkOpp7QPOqY6pdCP%2FGkxkRSFSTRm9l11%2B%2BsN7P%2F6nix%2B86Wf0Iuy6aUKjaZFYXItfDiBXI6yKLJ2wHaEQ9s2%2F%2BGz%2BVz7%2BwUZWya4rA9I9bOKqn6QO2uIeTT4TTXS9wQXo55IgnrMd02uKGLl7wMwt3YQ%2BuFa1DK4Mqy0Zd0drXE2i85kOpECOu2F7uBCxB3fGMmQRjTy60SVDabaXkjyndOyjqdoAHBl%2BtzZDKoYjya%2Foa6%2B04hvtz9OJAA0%2FCy3ci%2BJQ%2B%2FwC9frhhuj7tKXHskGXCrT61lMD3tV48FqiwVcTI9dEAtr7xlMD25UY2FrD5UIy0JoJQFvffGog%2B3IrZNxNJSYSweSnIED27acGnRhvH82u5pKCFH4cGnnnmLBqVKdOHC2fHk0kCSSVQOxnG9xx5WIk4vsriYKYH9wa%2B%2BVmfCJfNQQfERQOydMTl1B6ZZ7e66bGN92faa8UNu7PKciE%2BGYiTBoZNrT0blRv%2BVCNRVZdznAlI%2FLyTP7wMpLI%2Bhu0i%2B6IEs72TObzMfXktKpLblkwsV3cBGJ%2Fd5j08PVj%2BPrHo8j%2FxRb5i1xzbeSfO5ZhH2OAmVjwKSD2n%2BEOER0SwdRaqFilXL9DGSrjmvi2napuwjieIIPaKMgKbaRh7oYcn%2BZplRpQuJgp1K3Zo%2FmCxGbQJ82gz1sbk%2Fd%2Bx7xfr68UJrzf6e7sZBpa%2B1EbOwlbwxHs41SozBIeN%2Fal6i%2Fc1XHBkjZ%2FGZpcNSYxKCTFYBSK15mBw0Hgys%2BUqanRXZ07Bl0nZb5435GUbnDdoiu6gTehGmrcI0wjQV%2FrDbnWtoPrOM%2FRbb9JGmDPM4cXq16kEGPvEjGKJkCbHMpRgPIrUSh%2F4gUSf57iSN3lRIypHG%2Ba8ZXWNMM3sbFEOLhPARAvxgPRMa%2BIRt3AiMLx2f807gtue96LaoOyuD4zOpGflX1zgEHZk%2BkahZAl2LMOCegp2l6WKydd%2F0u0qAKyEQW83KoArYYuzSPTMYzInV%2BKMVit1o29wtxM5qUU7DVNvICnr6WD1Bfj7uTKSCKgtBpwdKETSsLCXIubzh9XPMzKCJD1HP64v2a03q6jZ77WpqNfa8XEm5Ggm%2FS2Ka7KuYluzGkM2h08SiQ51zZidIixkYgRZ3uJQe2RoPb4i4%2F9BcP%2BIt8uAyrMHGb0w2FouLeeGPfWInFPiv2uK7kmQK6HLf2AK%2F9xzSzZ1izq8HP4HUf4eXKINKLXMTvIrhQFw%2BLMBzGKurBJYVbv0SdIN2FFChd%2FTUQM1LSW4skTCUZ0mQ9bRIKf80825zAky9xvVUbm63cww54iSjqzyQ2EK5AoIMfESNU5Rv589%2BU9zi17ZmRE3oCrMDCjml4RmTgzqbu4DFaZqXTkjoEPCMth26q%2FwYjqMJvOb%2BIaqciK67K48%2FLCnLOAHziPPp0bcRZg0hvBNDYS4xZDf4gZ2IPFqEpsikQ0QPjA0TCSbKGyYwIZGATDGQOj5bXFzRuICjkqUBBHGE76fKOAcW%2FpW5jTt0hp2%2BH3dMYdbKxJGCwgLvmzHn1S0QUM%2B6QUcHngsqiSMIdiBEI0VNXtKs0l7pWJA8agDBiD%2FlaI6PYFf8tByzYBf1V%2F60p8f%2FV804p3f5V608ryLuMRjHIzpw6PR3DbK0cZt8GVRs1jMMUe1lU%2B%2BPWYjKahFT8xRNeuoZgxjkR2iVYtgwhvOB%2BTRL4ak25L%2BkXbU3tilRTkmPvTFqfwhCfexwz0xY7kuTKVCObQoyDY7x0J7MwTDXYThh3xDismDuzLMVbtuqAiDpxTcsOWFZDffwoMO02sY%2BYEo%2BTfyb%2FLldlEYEcjDEj4B88sPIp6jIV74fexTFwWTf7wmYknNPHRd5NOVTblSiDiHz2z8Qh51N6Lm14mmdzAZRnns2cGntDAC%2B8mXXkVTvlBvj9%2BZt2HWrccD3nGbXdj3bI69yfPrDuhdY8ltu7wyBXk%2B9MUrPvxJk4Lk%2F5GZX9zcbAhuLuhdUGWRE9HBTt8aJ2VDxxhaH3M06JpTrVcjOmrZaIWsQrlnW6whSpm9r3PW0dNgFjrnQKwPbzYNWJQJM3iZsTcRrlyNbpG3KwRxnPLmJNdEYWW3P%2BENbDiygmNlLFoBiZsNG14Y2P4QHzQ6upMtTpjA93sL8JyKU0QecLVErOalZM6p%2B7fWP28GGPRrrPoRjdQBs6WPm%2FW3AVsMWYdeiTpagIXcMBO%2FamHj7G8%2F%2FoE%2F5UH%2FmsKuowesjD7XNLocaTtkf8PC2razTBn3m61AF0j3h9%2Fr5jIQkYLY6P5sbFp9OGHMSlLh0fG8%2BKRQRTzGRhIRmg0eduFtNoeT9w2zBMktrIxqDVnW4dzXMFBCyLk6DU6g5YoE8maBds3YSKH2mhzbRc9cAiyCSKm6hBYHGCiml5xlwYSE6Z7SlhcFo%2FIhYKheaHcT9O5dVwEMZII9%2BnJbnCHUmi2dUjHM2bQwo5jW8TUdK0t9qpz4M8RWqr%2B6DMzognECex0xEYV1IFsnxy%2BgcRpBLuJTXTf0XSZ%2BFOENW%2BNJxb6hjk91YANrFjejR1OmXhWQ7CBjHBJkgoOsClcfZXA3J04I3CAqUnBZQ120%2BoqDvgCUGH3mLQYYQOazmBrFlzVTQDKb4oYjQlokyJSFzxBW3BsiztsHXZQOyVBnBn%2BHLVGq0IOQ6cBMZpDry%2BKJAaNTgh%2BDLRMVAZDocL4ZB4YLORHx64%2FLUYIs4HJjRCKldnW0S5XykELaAmbKjFweyNsdQC6iaXL%2BPx6gFKyIDWV2FlPT3SjL6i9ZVtHyjxDQRGbFDESXrqxIafVm%2FRmSKVp7jpu1fBWcYMSTejOYkCvE0ZNObkv9ERLhl7xOrF4KLQgobE4RFwwQTqPqGXgh6B7W77LQBX%2By%2FbuDbrgzakpuBtXqM2p2xnz4%2BhAdn%2FRpWFztuzvUbtyR4v6Q4GInUNIYOOZoLjTohp79InZWKdg6mYFlkrIh1z%2BpYcSMjUcnXue%2B1vSMdrDgmuJhoQCuziAKGEkkhtuu9zycReQjMkXUfx9myxR8bLEYMVDpwwRNnS3vgugueiQytb45GJCKUX5h6iYK8Bz39xIiSXYxt%2F9jhOZzE%2FKZH48Ef%2FeXnYg%2BHOHp%2FNHUdSA6M3CEgV9K%2FX3twSDmsabuzzZb%2FLw27F46IVWXYDibeuHVotJQGlXHROne84eNsa5GGOyHd6DdNqtci7KDeRceUkM%2BXPLS1VNnACBfK%2B%2B7fDYACNg06u4QtrEmNtQQZA3CBxyDjOKuGRTQzii4OxHVz6I3OQeFCV2xZKnm8lkuPtKBcBCmxGat8O771gLdgP5ZaLLLTv8Nwjfo1poi39DA8JeICi6jUdkP9sSLAXtV6980BpDGwQ%2BCq3p50q%2FK4UV3YcHneMlEWIcZgYJKFty%2B9%2Fr12ebSqnhp%2Fptp1TVeUffkch8AN%2Br8HWtXXGl53o0BYkWV1ZazIOJNAwy86OZh%2BF20wjNNzuVH%2BNzmLjV36gbDw7lx57opo%2BeXIsimtXyZU8MoTxPrWVs2KS9JoHsVEctvtyNFpXzybdbNnwlUAltpIA3rIX7WugtgU0dDh7MhzvUiTi1xvaml%2BN6k%2Ff0R02gxXxgnysp80jZKDU0RZ%2BwkjquINjA89H9Fj1b7eBK4H%2Fg2YnDcbEORSRxNvRy%2BxcFjcktm1ejAVLmPSuxb0mKe0mSONfycjA4Ccl5u8TJ%2Bj%2Fa9Db5z2YAAA%3D%3D&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT=vistaConsultaEstadoRUT%3AformConsultaEstadoRUT&vistaConsultaEstadoRUT%3AformConsultaEstadoRUT%3A_idcl=")

	req, err := http.NewRequest(http.MethodPost, url, data)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
	var razonSocial, estado, primerApellido, segundoApellido, primerNombre, otrosNombres, dv string

	var table = doc.Find("table.muisca_area")

	if table.Length() == 0 {
		return nil, ErrNoResults
	}

	dv = table.Find("#vistaConsultaEstadoRUT\\:formConsultaEstadoRUT\\:dv").Text()

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
		DV:              dv,
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
