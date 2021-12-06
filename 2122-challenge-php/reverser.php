<?php
function reverse(string $yo): string {
    return strrev($yo);
}
function isPalindrome(string $yo): bool {
    if (strrev($yo) == $yo){ 
        return 1; 
    }
    else{
        return 0;
    }
}