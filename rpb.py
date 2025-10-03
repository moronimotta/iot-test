from picozero import pico_led
import machine
import time
import json
import urequests
import network

# Initialize:
# Your WiFi network credentials
ssid = "Person"
password = "ts3jxpjtndwdwg8"

# Server to Post data to (ADD HTTP:// PREFIX!)
server_url = "http://172.31.208.1:8080/data"  # Fixed URL with http:// and port

#Pin for reading temperature
adcpin = 4
sensor = machine.ADC(adcpin)

# Functions
def ReadTemperature():
    adc_value = sensor.read_u16()
    volt = (3.3/65535) * adc_value
    c_temperature = 27 - (volt - 0.706)/0.001721
    f_temperature = (c_temperature * 9 / 5) + 32
    return round(f_temperature, 1)

def CreateJSON(temperature, humidity=50.0, device_id="pico_001"):
    # Data packet with multiple values
    data = {
        "device_id": device_id,
        "temperature": temperature,
        "humidity": humidity,
        "timestamp": time.time(),
        "status": "active"
    }
    
    json_packet = json.dumps(data)
    return json_packet

def ConnectToInternet(ssid, password):
    # Create a wireless interface object
    wlan = network.WLAN(network.STA_IF)
    wlan.active(True)

    # Connect to the network
    wlan.connect(ssid, password)
    
    # Wait for the connection to establish
    max_wait = 10
    while max_wait > 0:
        if wlan.status() < 0 or wlan.status() >= 3:
            break
        max_wait -= 1
        print('Waiting for connection...')
        time.sleep(1)
    
    # Check connection status
    if wlan.status() != 3:
        print('Network connection failed')
        return False
    else:
        print('Connected to WiFi')
        status = wlan.ifconfig()
        print('IP address:', status[0])
        return True

def SendData(json_data):
    try:
        headers = {'Content-Type': 'application/json'}
        response = urequests.post(server_url, data=json_data, headers=headers)
        print(f"Response status: {response.status_code}")
        print(f"Response text: {response.text}")
        response.close()
        return True
    except Exception as e:
        print(f"Error sending data: {e}")
        return False

def SendDummyData():
    """Send dummy JSON data for testing"""
    dummy_data = {
        "device_id": "pico_test",
        "temperature": 23.5,
        "humidity": 45.2,
        "light_level": 85,
        "timestamp": time.time(),
        "status": "testing"
    }
    
    json_data = json.dumps(dummy_data)
    print(f"Sending dummy data: {json_data}")
    return SendData(json_data)

# Main execution
def main():
    # Connect to WiFi
    if ConnectToInternet(ssid, password):
        pico_led.on()  # LED on when connected
        
        # Send dummy data every 10 seconds
        while True:
            try:
                print("Sending dummy data...")
                if SendDummyData():
                    print("Data sent successfully!")
                    pico_led.blink(n=2, on_time=0.2, off_time=0.2)
                else:
                    print("Failed to send data")
                    pico_led.blink(n=5, on_time=0.1, off_time=0.1)
                
                time.sleep(10)  # Wait 10 seconds before next send
                
            except KeyboardInterrupt:
                print("Stopping...")
                break
            except Exception as e:
                print(f"Error in main loop: {e}")
                time.sleep(5)
    else:
        print("Failed to connect to WiFi")
        # Blink LED rapidly to indicate error
        for _ in range(10):
            pico_led.on()
            time.sleep(0.1)
            pico_led.off()
            time.sleep(0.1)

# Run the main function
if __name__ == "__main__":
    main()