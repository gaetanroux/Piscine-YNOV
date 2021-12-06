<?php
function capsMe(string $yo): string {
    return strtoupper($yo);
}
function lowerMe(string $yo): string {
    return strtolower($yo);
}
function upperCaseFirst(string $yo): string {
    return ucwords($yo);
}
function lowerCaseFirst(string $yo): string {
    $test = explode(" ",$yo);
    $array = [];
    foreach ($test as $fin) {
    array_push($array, lcfirst($fin));
    }
    $res = implode(" ",$array);
    return $res;
}
function removeBlankSpace(string $yo): string {
    return trim($yo);
}