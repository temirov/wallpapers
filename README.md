# Wallpapers

This project is a self-contained webpage that generates grey wallpapers in two modes:

- **Tone Perfect Grey:** A uniform, flat grey image.
- **Gaussian Noise Grey:** A grey image with Gaussian noise (subtle variation) applied.

The page automatically detects your screen resolution but lets you override it if needed. It also includes SEO meta tags
and structured data (JSON‑LD) for better search engine discovery.

## Features

- **Two Modes:** Choose between a solid tone perfect grey and a Gaussian noise grey wallpaper.
- **Resolution Detection:** Auto-detects the user’s screen resolution using `screen.width`, `screen.height`, and
  `devicePixelRatio`.
- **Download:** Generate the wallpaper on an HTML5 canvas and download it as a PNG file.
- **SEO Friendly:** Comes with meta tags and JSON‑LD structured data to improve discoverability.

## Getting Started

1. **Clone or download** this repository.
2. **Open** the `index.html` file in your web browser.
3. **Select the Mode:** Use the dropdown to choose between "Tone Perfect Grey" and "Gaussian Noise Grey".
4. **Override Resolution (Optional):** If needed, manually enter your desired dimensions.
5. **Generate & Download:** Click the **Generate Wallpaper** button, then use the **Download Wallpaper** button to save
   your image.

## Examples

### Example 1: Tone Perfect Grey Wallpaper (600×800)

- **Mode:** Tone Perfect Grey
- **Resolution:** 600×800
- **Result:** A uniform grey image where every pixel has a grey value of 128.

### Example 2: Gaussian Noise Grey Wallpaper (600×800)

- **Mode:** Gaussian Noise Grey
- **Resolution:** 600×800
- **Result:** A grey image with Gaussian noise added around the grey value of 128 (with a standard deviation of 20) to
  produce a subtle, randomized texture.

To try these examples, open the webpage, manually set the resolution to 600 (width) and 800 (height), select your
desired mode, then click **Generate Wallpaper**.

## File Overview

- **index.html:** The main webpage that includes HTML, CSS, and JavaScript. It:
    - Detects screen resolution.
    - Generates the wallpaper on a canvas element.
    - Provides download functionality.
    - Contains SEO meta tags and JSON‑LD structured data.

## License

This project is licensed under the [MIT License](LICENSE).

