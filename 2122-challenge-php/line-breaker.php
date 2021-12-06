<?php
function breakLines(string $yo): string {
    $newtext = wordwrap($yo, 15, "\n", true);
} 