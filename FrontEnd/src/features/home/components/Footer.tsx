"use client";

import Image from "next/image";
import Link from "next/link";
import { ArrowRight, ArrowUp, ArrowUpRight, Sparkles } from "lucide-react";
import {
  FOOTER_BRAND,
  FOOTER_NAV,
  FOOTER_SOCIAL,
  FOOTER_CONTACT,
  FOOTER_LEGAL,
} from "@/features/home/data/data";
import { FooterCurves } from "@/features/home/components/FooterCurves";
import { useFooterAnimation } from "@/features/home/hooks/useFooterAnimation";

function ColumnTitle({ children }: { children: React.ReactNode }) {
  return (
    <h3 className="mb-4 text-[11.5px] font-bold tracking-[0.14em] text-white/50 uppercase">
      {children}
    </h3>
  );
}

export function Footer() {
  const { footerRef, handleFooterNavigation } = useFooterAnimation();

  return (
    <footer
      ref={footerRef}
      className="bg-primary relative isolate overflow-hidden text-white"
    >
      <div
        aria-hidden="true"
        className="absolute inset-0 z-0 bg-[linear-gradient(155deg,#1f506d_0%,#173c52_55%,#0a172b_100%)]"
      />

      <FooterCurves />

      <div className="relative z-10 container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid gap-10 py-14 sm:grid-cols-2 lg:grid-cols-[1.5fr_0.75fr_0.8fr_1.15fr] lg:gap-12 lg:py-16">
          <div data-footer-column className="flex flex-col items-start">
            <div className="flex items-center gap-3 rounded-2xl bg-white px-4 py-3 shadow-lg shadow-black/5">
              <Image
                src={FOOTER_BRAND.logo.src}
                alt={FOOTER_BRAND.logo.alt}
                width={FOOTER_BRAND.logo.width}
                height={FOOTER_BRAND.logo.height}
                sizes="56px"
                className="h-12 w-auto object-contain"
              />
              <div className="flex flex-col gap-0.5 leading-tight">
                <span className="text-primary text-lg font-extrabold tracking-tight">
                  {FOOTER_BRAND.org}
                </span>
                <span className="text-primary/60 text-[11px] font-medium">
                  Learning Tracker
                </span>
              </div>
            </div>

            <p className="mt-6 max-w-[35ch] text-sm leading-[1.85] text-white/65">
              {FOOTER_BRAND.tagline}
            </p>
          </div>

          <div data-footer-column>
            <ColumnTitle>Jelajahi</ColumnTitle>

            <ul className="flex flex-col gap-3">
              {FOOTER_NAV.map((link) => (
                <li key={link.label}>
                  <Link
                    href={link.href}
                    onClick={(e) => handleFooterNavigation(e, link.href)}
                    className="group inline-flex items-center gap-2 text-sm font-medium text-white/65 transition-colors hover:text-white focus-visible:rounded-sm focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-white"
                  >
                    <span>{link.label}</span>

                    <ArrowUpRight className="size-3.5 -translate-x-1 opacity-0 transition-all duration-200 group-hover:translate-x-0 group-hover:opacity-100" />
                  </Link>
                </li>
              ))}
            </ul>
          </div>

          <div data-footer-column>
            <ColumnTitle>Terhubung</ColumnTitle>

            <ul className="flex flex-col gap-3">
              {FOOTER_SOCIAL.map((link) => (
                <li key={link.label}>
                  <a
                    href={link.href}
                    target={link.external ? "_blank" : undefined}
                    rel={link.external ? "noopener noreferrer" : undefined}
                    className="group inline-flex items-center gap-2 text-sm font-medium text-white/65 transition-colors hover:text-white focus-visible:rounded-sm focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-white"
                  >
                    <span>{link.label}</span>

                    <ArrowUpRight className="size-3.5 -translate-x-1 opacity-0 transition-all duration-200 group-hover:translate-x-0 group-hover:opacity-100" />
                  </a>
                </li>
              ))}
            </ul>
          </div>

          <div data-footer-column>
            <ColumnTitle>{FOOTER_CONTACT.addressLabel}</ColumnTitle>

            <p className="max-w-[36ch] text-sm leading-[1.85] text-white/65">
              {FOOTER_CONTACT.address}
            </p>

            <div className="mt-7">
              <ColumnTitle>{FOOTER_CONTACT.contactLabel}</ColumnTitle>

              <a
                href={`mailto:${FOOTER_CONTACT.email}`}
                className="inline-flex max-w-full items-center gap-2 text-sm font-semibold break-all text-white/90 underline decoration-white/30 underline-offset-4 transition-colors hover:decoration-white"
              >
                {FOOTER_CONTACT.email}
                <ArrowUpRight className="size-3.5 shrink-0" />
              </a>
            </div>
          </div>
        </div>
      </div>
    </footer>
  );
}
