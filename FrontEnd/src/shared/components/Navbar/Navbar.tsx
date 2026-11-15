"use client";

import Link from "next/link";
import { Menu, X } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { useNavbar } from "@/shared/hooks/use-navbar";
import { Logo } from "../Logo/Logo";
import { NAV_LINKS } from "@/features/home/data/data";
import { cn } from "cn";

export function Navbar() {
  const { scrolled, menuOpen, toggleMenu, closeMenu } = useNavbar();

  return (
    <header className={cn("navbar", scrolled ? "scrolled" : "not-scrolled")}>
      <div className="navbar-inner flex w-full items-center justify-between gap-6">
        <Logo onNavigate={closeMenu} />

        <nav
          aria-label="Navigasi utama"
          className="hidden flex-1 items-center justify-center gap-2 lg:flex"
        >
          {NAV_LINKS.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              className="text-secondary-text hover:bg-secondary-light hover:text-primary focus-visible:ring-ring rounded-full px-4 py-2 text-base font-semibold transition-colors focus-visible:ring-2 focus-visible:outline-none"
            >
              {l.label}
            </Link>
          ))}
        </nav>

        <div className="hidden shrink-0 items-center gap-3 lg:flex">
          <Button
            variant="outline"
            className="border-primary/30 text-primary hover:border-primary hover:bg-secondary-light h-11 rounded-full px-6 text-base font-semibold"
          >
            <Link href="/auth">Masuk</Link>
          </Button>
          <Button className="group bg-primary hover:bg-primary/90 h-11 rounded-full px-7 text-base font-semibold text-white transition-all hover:-translate-y-0.5">
            <Link href="/auth">Mulai gratis</Link>
          </Button>
        </div>

        <button
          type="button"
          onClick={toggleMenu}
          aria-label={menuOpen ? "Tutup menu" : "Buka menu"}
          aria-expanded={menuOpen}
          aria-controls="mobile-menu"
          className="text-primary hover:bg-secondary-light ml-auto grid size-11 place-items-center rounded-full transition-colors lg:hidden"
        >
          {menuOpen ? <X size={24} /> : <Menu size={24} />}
        </button>
      </div>

      {menuOpen && (
        <nav
          id="mobile-menu"
          aria-label="Navigasi mobile"
          className="border-primary/15 bg-background animate-in fade-in slide-in-from-top-2 absolute inset-x-0 top-[calc(100%+10px)] flex flex-col gap-1 rounded-2xl border p-3 shadow-xl duration-200 lg:hidden"
        >
          {NAV_LINKS.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              onClick={closeMenu}
              className="text-secondary-text hover:bg-secondary-light hover:text-primary rounded-xl px-4 py-3 text-base font-semibold transition-colors"
            >
              {l.label}
            </Link>
          ))}

          <div className="border-border mt-1 flex flex-col gap-2 border-t pt-3">
            <Button
              variant="outline"
              className="border-primary/30 text-primary h-11 w-full rounded-xl text-base font-semibold"
            >
              <Link href="/auth" onClick={closeMenu}>
                Masuk
              </Link>
            </Button>
            <Button className="bg-primary hover:bg-primary/90 h-11 w-full rounded-xl text-base font-semibold text-white">
              <Link href="/auth" onClick={closeMenu}>
                Mulai gratis
              </Link>
            </Button>
          </div>
        </nav>
      )}
    </header>
  );
}
