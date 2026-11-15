import Image from "next/image"
import Link from "next/link"
import { ArrowUp, ArrowUpRight } from "lucide-react"
import {
  FOOTER_BRAND,
  FOOTER_NAV,
  FOOTER_SOCIAL,
  FOOTER_CONTACT,
  FOOTER_LEGAL,
} from "../Footer/footer-data"
import { FooterCurves } from "./FooterCurves"

function ColumnTitle({ children }: { children: React.ReactNode }) {
  return (
    <h3 className="mb-4 text-[11.5px] font-bold tracking-[0.14em] text-white/50 uppercase">
      {children}
    </h3>
  )
}

export function Footer() {
  return (
    <footer className="relative isolate overflow-hidden bg-primary text-white">
      <div
        aria-hidden
        className="absolute inset-0 z-0 bg-[linear-gradient(158deg,#1f506d_0%,#173c52_52%,#0a172b_100%)]"
      />

      <FooterCurves />

      <div
        aria-hidden
        className="absolute inset-x-0 top-0 z-10 h-px bg-linear-to-r from-transparent via-white/25 to-transparent"
      />

      <div className="relative z-10 mx-auto w-full max-w-7xl px-5 md:px-8 lg:px-10">
        <div className="grid gap-10 py-14 md:grid-cols-2 lg:grid-cols-[1.5fr_0.8fr_0.8fr_1.1fr] lg:gap-12">
          <div className="flex flex-col items-start">
            <div className="flex items-center gap-3 px-4 py-2 bg-white rounded-xl hover:bg-white/80">
              <Image
                src={FOOTER_BRAND.logo.src}
                alt={FOOTER_BRAND.logo.alt}
                width={FOOTER_BRAND.logo.width}
                height={FOOTER_BRAND.logo.height}
                sizes="56px"
                className="h-14 w-auto object-contain"
              />
              <div className="flex flex-col leading-tight">
                <span className="text-[15px] font-extrabold tracking-tight text-primary">
                  {FOOTER_BRAND.org}
                </span>
                <span className="text-[11.5px] font-medium text-primary/60">
                  {FOOTER_BRAND.orgLine2}
                </span>
              </div>
            </div>

            <p className="mt-5 max-w-[42ch] text-sm leading-relaxed text-white/70">
              {FOOTER_BRAND.tagline}
            </p>
          </div>

          <div>
            <ColumnTitle>Tautan</ColumnTitle>
            <ul className="space-y-2.5">
              {FOOTER_NAV.map((link) => (
                <li key={link.label}>
                  <Link
                    href={link.href}
                    className="text-sm font-semibold text-white/70 transition-colors hover:text-white focus-visible:ring-2 focus-visible:ring-white/40 focus-visible:outline-none"
                  >
                    {link.label}
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          <div>
            <ColumnTitle>Sosial Media</ColumnTitle>
            <ul className="space-y-2.5">
              {FOOTER_SOCIAL.map((link) => (
                <li key={link.label}>
                  <a
                    href={link.href}
                    target={link.external ? "_blank" : undefined}
                    rel={link.external ? "noopener noreferrer" : undefined}
                    className="group inline-flex items-center gap-1 text-sm font-semibold text-white/70 transition-colors hover:text-white"
                  >
                    {link.label}
                    <ArrowUpRight className="size-3 -translate-x-1 opacity-0 transition-all group-hover:translate-x-0 group-hover:opacity-100" />
                  </a>
                </li>
              ))}
            </ul>
          </div>

          <div>
            <ColumnTitle>{FOOTER_CONTACT.addressLabel}</ColumnTitle>
            <p className="max-w-[34ch] text-sm leading-relaxed text-white/70">
              {FOOTER_CONTACT.address}
            </p>

            <div className="mt-6">
              <ColumnTitle>{FOOTER_CONTACT.contactLabel}</ColumnTitle>
              <a
                href={`mailto:${FOOTER_CONTACT.email}`}
                className="text-sm font-semibold text-white underline decoration-white/30 underline-offset-4 transition-colors hover:decoration-white"
              >
                {FOOTER_CONTACT.email}
              </a>
            </div>
          </div>
        </div>

        <div className="flex flex-col gap-4 border-t border-white/15 py-6 sm:flex-row sm:items-center sm:justify-between">
          <p className="text-[13px] text-white/55">{FOOTER_LEGAL.copyright}</p>

          <div className="flex items-center gap-5">
            <span className="text-[13px] text-white/55">{FOOTER_LEGAL.note}</span>
            <Link
              href="#top"
              className="inline-flex items-center gap-1.5 rounded-full border border-white/20 bg-white/5 px-3.5 py-1.5 text-[13px] font-semibold text-white/80 backdrop-blur-sm transition-colors hover:border-white/40 hover:bg-white/15 hover:text-white"
            >
              <ArrowUp className="size-3.5" strokeWidth={2.2} />
              Ke atas
            </Link>
          </div>
        </div>
      </div>
    </footer>
  )
}