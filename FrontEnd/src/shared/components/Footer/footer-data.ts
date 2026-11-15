export type FooterLink = {
  label: string
  href: string
  external?: boolean
}

export const FOOTER_BRAND = {
  org: "Lentera",
  orgLine2: "Lentera - Learning Tracker",
  tagline:
    "Wadah belajar, berkarya, dan bertumbuh bersama. Satu tempat untuk belajar dan mengajarkan",
  logo: {
    src: "/logo.webp",
    alt: "Logo Lentera",
    width: 112,
    height: 146,
  },
} as const

export const FOOTER_NAV: FooterLink[] = [
  { label: "Beranda", href: "#top" },
  { label: "Cara Menggunakan", href: "#cara-menggunakan" },
  { label: "Manfaat", href: "#manfaat" },
  { label: "Tentang", href: "#tentang" },
]

export const FOOTER_SOCIAL: FooterLink[] = [
  { label: "Instagram", href: "https://www.instagram.com/", external: true },
  { label: "LinkedIn", href: "https://www.linkedin.com/", external: true },
  { label: "YouTube", href: "https://www.youtube.com/", external: true },
  { label: "Twitter", href: "https://x.com/", external: true },
  { label: "TikTok", href: "https://www.tiktok.com/", external: true },
]

export const FOOTER_CONTACT = {
  addressLabel: "Alamat",
  address:
    "Gedung GKM Fakultas Ilmu Komputer Universitas Brawijaya Lantai 2, Universitas Brawijaya, Jalan Veteran 06C, Malang 66145",
  contactLabel: "Kontak",
  email: "lentera@ub.ac.id",
} as const

export const FOOTER_LEGAL = {
  copyright: "© 2026 Lentera - Learning Tracker",
  note: "",
} as const