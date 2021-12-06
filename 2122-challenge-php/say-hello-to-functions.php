<?php
function sayHello(): string
{
    return "Hello";
}
function sayHelloTo(string $name): string
{
    return "Hello $name";
}