package kvm

import (
	"encoding/json"
	"fmt"
	"os"
)

type WakeOnLanDevice struct {
	Name       string `json:"name"`
	MacAddress string `json:"macAddress"`
}

type UsbConfig struct {
	VendorId     string `json:"vendor_id"`
	ProductId    string `json:"product_id"`
	SerialNumber string `json:"serial_number"`
	Manufacturer string `json:"manufacturer"`
	Product      string `json:"product"`
}

type Config struct {
	CloudURL             string            `json:"cloud_url"`
	CloudToken           string            `json:"cloud_token"`
	GoogleIdentity       string            `json:"google_identity"`
	JigglerEnabled       bool              `json:"jiggler_enabled"`
	AutoUpdateEnabled    bool              `json:"auto_update_enabled"`
	IncludePreRelease    bool              `json:"include_pre_release"`
	HashedPassword       string            `json:"hashed_password"`
	LocalAuthToken       string            `json:"local_auth_token"`
	LocalAuthMode        string            `json:"localAuthMode"` //TODO: fix it with migration
	WakeOnLanDevices     []WakeOnLanDevice `json:"wake_on_lan_devices"`
	DisplayMaxBrightness int               `json:"display_max_brightness"`
	DisplayDimAfterSec   int               `json:"display_dim_after_sec"`
	DisplayOffAfterSec   int               `json:"display_off_after_sec"`
	EdidString           string            `json:"hdmi_edid_string"`
	UsbConfig            UsbConfig         `json:"usb_config"`
	VirtualMediaEnabled  bool              `json:"virtual_media_enabled"`
}

const configPath = "/userdata/kvm_config.json"

var defaultConfig = &Config{
	CloudURL:             "https://api.jetkvm.com",
	AutoUpdateEnabled:    true, // Set a default value
	DisplayMaxBrightness: 64,
	DisplayDimAfterSec:   120,  // 2 minutes
	DisplayOffAfterSec:   1800, // 30 minutes
	VirtualMediaEnabled:  true,
	UsbConfig: UsbConfig{
		VendorId:     "0x1d6b", //The Linux Foundation
		ProductId:    "0x0104", //Multifunction Composite Gadget
		SerialNumber: "",
		Manufacturer: "JetKVM",
		Product:      "JetKVM USB Emulation Device",
	},
}

var config *Config

func LoadConfig() {
	if config != nil {
		return
	}

	file, err := os.Open(configPath)
	if err != nil {
		logger.Debug("default config file doesn't exist, using default")
		config = defaultConfig
		return
	}
	defer file.Close()

	var loadedConfig Config
	if err := json.NewDecoder(file).Decode(&loadedConfig); err != nil {
		logger.Errorf("config file JSON parsing failed, %v", err)
		config = defaultConfig
		return
	}

	config = &loadedConfig
}

func SaveConfig() error {
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	return nil
}
