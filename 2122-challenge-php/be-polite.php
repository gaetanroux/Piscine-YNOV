<?php

date_default_timezone_set('Europe/Paris');
$hour = date('h');

if ($hour >6 && $hour <12 ) {
    echo "Hello! Have a nice day :)";
} elseif ($hour >12 && $hour <6) {
    echo "Have a good afternoon!";
} elseif ($hour >6 && $hour <9) {
    echo "Good evening! Hope you had a good day!";
} else {
    echo "Good night! See you tomorrow :)";

}