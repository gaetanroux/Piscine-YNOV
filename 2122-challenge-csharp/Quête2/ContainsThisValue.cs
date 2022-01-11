namespace CSharpDiscovery.Quest02
{
    public class ContainsThisValue_Exercice
    {
        public static bool ContainsThisValue(int[] tab, int value)
        {

            if (tab == null) {
                return false;
            }

            foreach (int i in tab) {
                if (i == value) {
                    return true;
                }
            }
            
            return false;
        }
    }
}