namespace CSharpDiscovery.Quest02
{
    public class FindLastIndex_Exercice
    {
        public static int? FindLastIndex(int[] tab, int a)
        {
            int res = -1;

            if (tab == null) {
                return null;
            }
            
            for (int i = 0; i < tab.Length; i++) {
                if (tab[i] == a) {
                    res = i;
                }
            }
            
            if (res == -1) {
                return null;
            }

            return res;
        }
    }
}