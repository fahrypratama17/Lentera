"use client"

import Link from "next/link"
import { Menu, X } from "lucide-react"
import { Button } from "@/shared/components/ui/button"
import { useNavbar } from "@/shared/hooks/use-navbar"
import { Logo } from "../Logo/Logo"
import { NAV_LINKS } from "@/features/home/data/data"
import { cn } from "cn"

export function Navbar() {
  const { scrolled, menuOpen, toggleMenu, closeMenu } = useNavbar()

  return (
    <header className={cn("navbar", scrolled ? "scrolled" : "not-scrolled")}>
      <div className="navbar__inner">
        <Logo onNavigate={closeMenu} />

        <nav
          aria-label="Navigasi utama"
          className="hidden items-center gap-1 lg:flex"
        >
          {NAV_LINKS.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              className="rounded-full px-3.5 py-2 text-sm font-semibold text-secondary-text
                transition-colors hover:bg-secondary-light hover:text-primary
                focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
            >
              {l.label}
            </Link>
          ))}
        </nav>

        <div className="hidden shrink-0 items-center gap-2 lg:flex">
          <Button
            variant="outline"
            className="h-10 cursor-pointer rounded-full border-primary/30 bg-transparent px-5
              text-sm font-semibold text-primary
              hover:border-primary hover:bg-secondary-light"
          >
            <Link href="/auth">Masuk</Link>
          </Button>

          <Button
            className="group h-10 cursor-pointer rounded-full bg-primary px-6 text-sm font-semibold
              text-white shadow-[0_1px_2px_rgba(10,23,43,0.06)]
              transition-all hover:-translate-y-0.5 hover:bg-primary/90
              hover:shadow-[0_10px_24px_-12px_rgba(31,80,109,0.7)] cursor-pointer"
          >
            <Link href="/auth">Mulai gratis</Link>
          </Button>
        </div>

        <button
          type="button"
          onClick={toggleMenu}
          aria-label={menuOpen ? "Tutup menu" : "Buka menu"}
          aria-expanded={menuOpen}
          aria-controls="mobile-menu"
          className="ml-auto grid size-10 place-items-center rounded-full text-primary
            transition-colors hover:bg-secondary-light lg:hidden"
        >
          {menuOpen ? <X size={22} /> : <Menu size={22} />}
        </button>
      </div>

      {menuOpen && (
        <nav
          id="mobile-menu"
          aria-label="Navigasi mobile"
          className="absolute inset-x-0 top-[calc(100%+10px)]
            flex flex-col gap-1 rounded-2xl border border-primary/15
            bg-background p-3 shadow-[0_18px_40px_-18px_rgba(10,23,43,0.35)]
            duration-200 animate-in fade-in slide-in-from-top-2 lg:hidden"
        >
          {NAV_LINKS.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              onClick={closeMenu}
              className="rounded-xl px-4 py-3 text-sm font-semibold text-secondary-text
                transition-colors hover:bg-secondary-light hover:text-primary"
            >
              {l.label}
            </Link>
          ))}

          <div className="mt-1 flex flex-col gap-2 border-t border-border pt-3">
            <Button
              variant="outline"
              className="h-11 w-full rounded-xl border-primary/30 text-sm font-semibold text-primary"
            >
              <Link href="/login" onClick={closeMenu}>
                Masuk
              </Link>
            </Button>

            <Button
              className="h-11 w-full rounded-xl bg-primary text-sm font-semibold text-white hover:bg-primary/90"
            >
              <Link href="/register" onClick={closeMenu}>
                Mulai gratis
              </Link>
            </Button>
          </div>
        </nav>
      )}
    </header>
  )
}