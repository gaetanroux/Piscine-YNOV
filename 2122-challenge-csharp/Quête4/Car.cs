namespace CSharpDiscovery.Quest04
{
    public class Car : Vehicule
    {
        public string Model {get; set;}

        public Car() : base() {
            Model = "Unknown";
        }

        public Car(string Model, string Brand, string Color, int CurrentSpeed = 0) : base(Brand, Color, CurrentSpeed) {
            this.Model = Model;
        }

        public override string ToString()
        {
            return this.Color + " " + this.Brand + " " + this.Model;
        }

        public override void Accelerate(int Speed)
        {
            this.CurrentSpeed += Speed;
            if (this.CurrentSpeed > 180) {
                this.CurrentSpeed = 180;
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