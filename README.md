# Wallpapers

This project is a self‑contained webpage that generates minimalist wallpapers in four modes:

- **Tone Perfect Grey:** A uniform, flat grey image.
- **Gaussian Noise Grey:** A grey image with subtle Gaussian noise.
- **Tone Perfect Pink:** A uniform, flat pink image.
- **Pink Noise:** A pink‑tone image with subtle pink noise texture.

The page automatically detects your screen’s native resolution and displays live previews. Click any card’s **Download
Full Resolution** button to instantly generate and save a PNG at your display’s exact dimensions. Built with Bootstrap 5
for a responsive, mobile‑friendly interface, and includes full SEO meta tags + JSON‑LD structured data.

## Features

- **Four Modes:** Tone‑perfect grey, Gaussian noise grey, tone‑perfect pink, pink noise
- **Automatic Resolution Detection:** Uses `screen.width`, `screen.height` & `devicePixelRatio`
- **Live Previews:** Low‑res thumbnails update instantly
- **One‑Click Download:** Full‑resolution PNG generated client‑side
- **Responsive Design:** Bootstrap 5 layout adapts to any device
- **SEO Friendly:** Meta tags + JSON‑LD for discoverability

## Getting Started

1. Clone or download this repository.
2. Open `index.html` in your browser.
3. View previews and click **Download Full Resolution** on any card to save your wallpaper.

## Examples (600×800)

| Mode                | Description                            |
|---------------------|----------------------------------------|
| Tone Perfect Grey   | Uniform grey (#808080)                 |
| Gaussian Noise Grey | Grey (#808080) + Gaussian noise (σ=20) |
| Tone Perfect Pink   | Uniform pink (#FFC0CB)                 |
| Pink Noise          | Pink (#FFC0CB) + pink noise texture    |

To reproduce any example, simply resize your browser window (or override resolution in dev tools) to 600×800, then click
**Download Full Resolution**.

## File Overview

- **index.html** — All HTML, CSS (Bootstrap), JavaScript, SEO meta tags, and image‑generation logic live here.

## License

This project is licensed under the MIT License.  
