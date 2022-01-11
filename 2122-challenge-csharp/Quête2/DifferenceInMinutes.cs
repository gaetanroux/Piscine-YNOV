using System;

namespace CSharpDiscovery.Quest02
{
    public class DifferenceInMinutes_Exercice
    {
        public static double DifferenceInMinutes(DateTime start, DateTime end)
        {
            TimeSpan Diff_dates = end.Subtract(start); 
            return Diff_dates.TotalDays * 24 * 60;
        }
    }
}