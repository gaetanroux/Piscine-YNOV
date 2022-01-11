using System;
using System.Globalization;

namespace CSharpDiscovery.Quest02
{
    public class FromStringToDateTime_Exercice
    {
        public static DateTime FromStringToDateTime(string dateStr)
        {
            return DateTime.Parse(dateStr, CultureInfo.GetCultureInfo("fr-FR"));
        }
    }
}