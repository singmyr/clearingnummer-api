<?php

require 'vendor/autoload.php';

header('Content-Type: application/json');

$output = [];

$host = 'https://www.swedishbankers.se';
try {
    $page = file_get_contents($host . '/fraagor-vi-arbetar-med/clearingnummer/clearingnummer');
    if (preg_match_all('/<a[^>]+href=\"([^\"]+)\"[^>]*><span>Clearingnummer - /', $page, $matches)) {
        $pdf = $host . $matches[1][0];
    } else {
        echo "shit";
        die();
    }

    if (isset($pdf) && $pdf) {
        $parser = new \Smalot\PdfParser\Parser();
        $pdf = $parser->parseFile($pdf);

        foreach ($pdf->getPages() as $page) {
            // echo $page->getText();
            if (preg_match_all('/^([a-zA-ZåäöÅÄÖ]+[a-zA-ZåäöÅÄÖ\/\&\(\)\- ]+)(\d+-{0,1}\d*)/m', $page->getText(), $matches)) {
                if (isset($matches[1]) && isset($matches[2])) {
                    foreach ($matches[1] as $i => $bank) {
                        $output[trim($bank)][] = $matches[2][$i];
                    }
                }
            }
        }
    }
} catch (\Exception $e) {
    echo $e->getMessage();
}

echo json_encode($output);