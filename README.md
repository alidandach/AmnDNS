# AmnDNS – Secure, AI-Powered DNS Server for Blocking Ads, Trackers, and Malicious Traffic

**AmnDNS** (أمن) is an open-source, AI-powered DNS server designed to provide users with enhanced security and privacy by blocking advertisements, user tracking, and malicious activities. The name "Amn" means "security" in Arabic, reflecting the core purpose of this project: to safeguard users from unwanted threats on the internet.

## Key Features:
- **Ad Blocking**: Automatically filters out advertisements from DNS queries, enhancing the browsing experience.
- **Tracking Protection**: Blocks domains used for user tracking and analytics to ensure better privacy.
- **AI-Driven Analysis**: Integrates AI to analyze DNS traffic patterns and detect emerging threats in real-time.
- **Customizable Blocklists**: Supports user-defined blocklists, allowing for tailored filtering of unwanted content.
- **Lightweight and Efficient**: Optimized for high performance with minimal resource usage, suitable for personal and enterprise environments.
- **Easy Integration**: Simple to set up and integrate into existing network infrastructures.

## Why AmnDNS?
In today's online environment, privacy and security are critical. AmnDNS not only blocks intrusive ads and trackers but also utilizes artificial intelligence to identify and block potential threats dynamically. Whether you're a personal user seeking a safer browsing experience or an organization looking to protect your network, AmnDNS offers a reliable, scalable solution.

## Project Structure:

The following is the suggested directory structure for **AmnDNS** to ensure scalability and maintainability.
```bash
AmnDNS/
│
├── cmd/               # Main application entry points
│   └── amndns/        # The DNS server main program
│       └── main.go    # Main application file
│
├── internal/          # Internal packages 
│   ├── dns/           # DNS-specific logic (resolvers, handling queries)
│   └── blocker/       # Ad/tracking/malicious domain blocking logic
│
├── pkg/               # Publicly available packages (if any) for external use
│
├── config/            # Configuration files (YAML, JSON, TOML)
│   └── config.yaml    # Example: Configuration for DNS server settings
│
├── docs/              # Documentation
│   └── architecture.md # Documentation for how the DNS server is structured
│
├── test/              # Unit and integration tests
│   └── dns_test.go    # Test file for DNS entry function
│
├── scripts/           # Helper scripts for automation (e.g., building, running)
│
├── LICENSE            # MIT License
├── README.md          # Project description
├── go.mod             # Go module dependencies
└── go.sum             # Dependency checksums (generated)
```

### Directory Descriptions:

- **cmd/amndns/**: Contains the main entry point for the DNS server.
- **internal/dns/**: Implements DNS query resolution, server setup, and request handling.
- **internal/blocker/**: Contains the logic for blocking ad, tracking, and malicious domains.
- **pkg/**: Public reusable packages, if any, for external use.
- **config/**: Stores configuration files, such as blocklist sources and server settings.
- **docs/**: Contains documentation related to architecture, design, or usage.
- **test/**: Holds unit and integration tests to validate the functionality of the DNS server.
- **scripts/**: Contains helper scripts to automate project-related tasks like build and deployment.


## Contribute:
AmnDNS is an open-source project, and contributions are welcome! Whether you want to suggest new features, report bugs, or contribute code, we'd love your help to improve and expand this project.

## License:
This project is licensed under the [MIT License](./LICENSE), making it free for personal and commercial use, with attribution.
