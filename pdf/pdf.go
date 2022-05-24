package pdf

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func Gen(url string) ([]byte, error) {
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cancelFunc := chromedp.NewExecAllocator(ctx, options...)
	defer cancelFunc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			printToPDFParams := page.PrintToPDF().WithPrintBackground(true)
			//{"marginTop":5e-324,"marginBottom":5e-324,"marginLeft":5e-324,"marginRight":5e-324}
			printToPDFParams.MarginTop = .4
			printToPDFParams.MarginBottom = .4
			printToPDFParams.MarginLeft = .5
			printToPDFParams.MarginRight = .5
			buf2, _, err := printToPDFParams.Do(ctx)
			if err != nil {
				return err
			}
			buf = buf2
			return nil
		}),
	); err != nil {
		return nil, err
	}
	//if err := ioutil.WriteFile("sample.pdf", buf, 0644); err != nil {
	//	log.Fatal(err)
	//}
	return buf, nil
}
