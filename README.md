# 🧲 BitTorrent Peer-to-Peer File Sharing Client

This project is an implementation of a **BitTorrent client**, a peer-to-peer (P2P) protocol for downloading and distributing files over the Internet. Unlike traditional client-server models—such as streaming a video on Netflix or loading a website—BitTorrent enables users (called **peers**) to download pieces of files **from each other**, making it a fully decentralized file-sharing system.

---

## 🔍 Overview

In this project, we explore the underlying mechanisms of the BitTorrent protocol and build a custom client capable of:

- Connecting to a BitTorrent tracker
- Discovering and connecting to multiple peers
- Requesting and exchanging file pieces between peers
- Assembling the complete file from downloaded chunks

---

## 💡 Key Concepts

- **Peer-to-Peer (P2P)**: Files are shared directly between users, without relying on a central server.
- **Tracker**: A server that helps clients find each other.
- **Torrent File**: Contains metadata about the file to be shared and the tracker URL.
- **Swarm**: The group of peers sharing a specific file.
- **Pieces and Blocks**: Files are split into chunks (pieces), which are further divided into smaller blocks.

---

## 🛠️ Features

- Parse `.torrent` files
- Communicate with trackers to get peer information
- Establish TCP connections with multiple peers
- Download file pieces concurrently from available peers
- Reconstruct the full file locally

## Project Structure

```
BITTORRENTCLIENT/
├── bitfield/                   # Logic for managing the BitField structure (which pieces a peer has)
├── client/                     # Core BitTorrent client logic
├── handshake/                  # Implements the BitTorrent handshake protocol
├── message/                    # Handles encoding and decoding of peer protocol messages
├── p2p/                        # Peer-to-peer communication logic
│   └── p2p.go
├── peers/                      # Structures and utilities for managing peer connections
│   └── peers.go
├── torrentfiles/               # Handles .torrent file parsing and tracker communication
│   ├── torrent.go             # Parses .torrent metadata
│   └── tracker.go             # Connects to tracker and retrieves peer list
├── BITTORRENTCLIENT           # Project metadata
├── debian-12.11.0-...         # Test/demo torrent file
├── go.mod                     # Go module definition (dependencies and module name)
├── go.sum                     # Checksums for dependencies
└── main.go                    # Entry point of the application
```

## Features

- **Torrent File Parsing**: Parse .torrent files and extract metadata
- **Tracker Communication**: Connect to trackers and retrieve peer lists
- **Peer-to-Peer Protocol**: Implement BitTorrent peer protocol for file sharing
- **Handshake Protocol**: Handle peer handshakes and connection establishment
- **Message Handling**: Encode/decode BitTorrent protocol messages
- **BitField Management**: Track which pieces each peer has available
- **File Download**: Download files through the BitTorrent protocol

## Getting Started

### Prerequisites

- Go 1.19+ installed on your system

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/bittorrent-client.git
   cd bittorrent-client
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Usage

```bash
go run main.go [torrent-file]
```

Example:
```bash
go run main.go debian-12.11.0-amd64-netinst.iso.torrent
```

## Components

### Core Modules

- **`bitfield/`**: Manages the bitfield data structure that tracks which pieces of a file each peer has
- **`client/`**: Contains the main client logic and coordination between components
- **`handshake/`**: Implements the initial handshake process when connecting to peers
- **`message/`**: Handles all BitTorrent protocol message types (interested, have, bitfield, etc.)
- **`p2p/`**: Manages peer-to-peer connections and communication
- **`peers/`**: Peer management utilities and data structures
- **`torrentfiles/`**: Torrent file parsing and tracker communication

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built following the BitTorrent Protocol Specification
- Inspired by the need to understand P2P file sharing protocols
