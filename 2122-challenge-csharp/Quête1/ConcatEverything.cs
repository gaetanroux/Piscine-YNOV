using System;
namespace CSharpDiscovery.Quest01
{
    public class ConcatEverything_Exercice
    {
        public static string ConcatEverything(params string[] str)
        {
            string newstr = "";
           for (int i = 0; i < str.Length; i++) {
               newstr += str[i];
           }
           return newstr;
        }
    }
}
