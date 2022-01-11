namespace CSharpDiscovery.Quest03
{
    public class HistoricalMonument : PointOfInterest
    {
        public int BuildYear {get; set;}

        public HistoricalMonument() {
            
        }

        public HistoricalMonument(string Name, double Latitude, double Longitude, int buildYear) : base(Name, Latitude, Longitude) {
            BuildYear = buildYear;
        }

        public override string ToString()
        {
            return this.Name + " is a historical monument built in " + this.BuildYear;
        }
    }
}