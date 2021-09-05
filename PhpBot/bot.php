<?php

$GET_INPUT = file_get_contents('php://input');

const TOKEN = '1958943268:AAFJkPyMNRN4cUVJhcVgwHIL3SnMMsH7Gn8';

const API_URL = 'https://api.telegram.org/bot';


function printAnswer($str) {
    echo "<pre>";
    print_r($str);
    echo "</pre>";
}


function getTelegramApi($method, $options = null) {

    $str_request = API_URL . TOKEN . '/' . $method;

    if ($options) {
        $str_request .= '?' . http_build_query($options);
    }

    $request = file_get_contents($str_request);

    return json_decode($request, 1);
}

function setHook($set = 1 ) {
    $url = "https://" . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI'];
    printAnswer(
        getTelegramApi('setWebhook'),
            [
                'url' => $set?$url:''
            ]
    )
    exit(); 
}

function createAnswer($poll_flag = 0, $name_surname = "", $event) {

    switch ($poll_flag) {
        case 0:
            if ($event['message']['text'] == '/start') {
                $poll_flag = 1;
                $respond_text = "Привет, введи свое имя";
            }
            break;
        case 1:
            $poll_flag = 2;
            $respond_text = "Привет, введи фамилию";
            $name_surname = $event['message']['text'];
            break;

        case 2:
            $poll_flag = 0;
            $respond_text = 'Очень приятно ' . $event['message']['text'] . ' ' . $name_surname;
            break;
    }

    return $poll_flag, $name_surname, $respond_text;
}


// установка веб хука, достаточно выполнить 1 раз
setHook(1);


$event = json_decode($GET_INPUT, 1);
//вызов функции и присваивание, хз точно как это тут работает 
$poll_flag, $name_surname, $respond_text = createAnswer($poll_flag, $name_surname, $event);

getTelegramApi('sendMessage', 
    [
        'text' => $respond_text,
        'chat_id' => $event['message']['chat']['id']
    ]
);

