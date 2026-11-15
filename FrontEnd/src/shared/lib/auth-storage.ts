export const AUTH_TOKEN_KEY = "lentera_token";

export function getAccessToken(): string | null {
  if (typeof window === "undefined") return null;

  return localStorage.getItem(AUTH_TOKEN_KEY);
}

export function setAccessToken(token: string): void {
  localStorage.setItem(AUTH_TOKEN_KEY, token);
}

export function removeAccessToken(): void {
  localStorage.removeItem(AUTH_TOKEN_KEY);
}
