<?php
for ($i = 2; $i < 96; $i++) {
    $isPrime = true;
    for ($p = 2; $p < $i; $p++) {
        if ($i % $p === 0) {
            $isPrime = false;
            break;
        }
    }
    if ($isPrime) {echo $i.', ';}
    
}   echo'97';
