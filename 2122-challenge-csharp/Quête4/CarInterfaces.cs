namespace CSharpDiscovery.Quest04
{
    public interface IThermalCar {
        public int FuelLevel {get; set;}

        void FillUp();

        int GetFuelLevel();
    }

    public interface IElectricCar {
        public int BatteryLevel {get; set;}

        void Recharge();

        int GetBatteryLevel();
    }
}