<div align="center">

  [![Unlicense License][license-shield]][license-url]
  [![LinkedIn][linkedin-shield]][linkedin-url]

</div>


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">GoLang Orderbook (Simulation)</h3>

  <p align="center">
    An GoLang Orderbook simulation
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

* Can create a orderbook for a single asset
  * Can fulfill a limit order, which is a collection of partial fill bids on asks
  * Price of asset should be impacts based on the supply and available orders
  * A MatchingEngine handles the price matching for the orderbook prices available using FIFO
  * When a Limit order is made, the collection of asks with their prices fulfilled, as well as the spread is output to the console
* Can create a OB Simulation
  * Create (x) asset, with initial supply (x) and value (x)
  * Creates (x) market participants who hold a random supply
  * Every (x) seconds, market participants place random Buys/Sells (bids/asks)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

![Go][go-shield]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Clone the repo
   ```sh
   https://github.com/luke-karmic/go-exchange
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Run the simulation
   ```sh
   go main.go
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Create Orderbook structure
- [ ] Create Limit, Asks, Bids structure
- [ ] Limit order market making
- [ ] Create market participants
- [ ] Create simulation

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Luke Taaffe

Project Link: [https://github.com/luke-karmic/go-exchange](https://github.com/luke-karmic/go-exchange)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/luketaaffe/
[go-shield]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge