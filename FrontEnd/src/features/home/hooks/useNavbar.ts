import type { MouseEvent } from "react";

export const handleSectionClick = (
  e: MouseEvent<HTMLAnchorElement>,
  href: string,
  closeMenu: () => void,
) => {
  if (!href.startsWith("#")) {
    closeMenu();
    return;
  }

  e.preventDefault();

  const section = document.getElementById(href.slice(1));

  if (section) {
    section.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });

    window.history.replaceState(null, "", href);
  }

  closeMenu();
};
