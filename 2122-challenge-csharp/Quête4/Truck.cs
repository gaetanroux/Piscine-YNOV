namespace CSharpDiscovery.Quest04
{
    public class Truck : Vehicule
    {
        public int Tonnage { get; set; } = 0;

        public Truck() : base() {
            Tonnage = 0;
        }

        public Truck(int Tonnage, string Brand, string Color, int CurrentSpeed = 0) : base(Brand, Color, CurrentSpeed) {
            this.Tonnage = Tonnage;
        }

        public override string ToString()
        {
            return this.Color + " " + this.Brand + " " + this.Tonnage + "T Truck";
        }

        public override void Accelerate(int Speed)
        {
            this.CurrentSpeed += Speed;
            if (this.CurrentSpeed > 100) {
                this.CurrentSpeed = 100;
            }
        }

        public override void Brake(int BrakePower)
        {
            this.CurrentSpeed -= BrakePower;
            if (this.CurrentSpeed < 0) {
                this.CurrentSpeed = 0;
            }
        }
    }
}