"use client";

import { useLayoutEffect, useRef } from "react";
import type { MouseEvent } from "react";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";

gsap.registerPlugin(ScrollTrigger);

export function useFooterAnimation() {
  const footerRef = useRef<HTMLElement>(null);

  const handleFooterNavigation = (
    e: MouseEvent<HTMLAnchorElement>,
    href: string,
  ) => {
    if (!href.startsWith("#")) return;

    const section = document.getElementById(href.slice(1));

    if (!section) return;

    e.preventDefault();

    section.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });

    window.history.replaceState(null, "", href);
  };

  useLayoutEffect(() => {
    const footer = footerRef.current;
    if (!footer) return;

    const mm = gsap.matchMedia();

    mm.add("(prefers-reduced-motion: no-preference)", () => {
      const ctx = gsap.context(() => {
        const columns = gsap.utils.toArray<HTMLElement>(
          "[data-footer-column]",
          footer,
        );

        gsap.fromTo(
          columns,
          {
            autoAlpha: 0,
            y: 35,
          },
          {
            autoAlpha: 1,
            y: 0,
            duration: 0.7,
            stagger: 0.12,
            ease: "power3.out",
            clearProps: "opacity,visibility,transform",
            scrollTrigger: {
              trigger: footer,
              start: "top 85%",
              once: true,
            },
          },
        );
      }, footer);

      return () => ctx.revert();
    });

    return () => mm.revert();
  }, []);

  return {
    footerRef,
    handleFooterNavigation,
  };
}
