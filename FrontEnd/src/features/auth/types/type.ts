import { Control, FieldValues, Path } from "react-hook-form";

export interface AuthHeaderProps {
  title: string;
  description: string;
}

export interface AuthInputProps<T extends FieldValues> {
  control: Control<T>;
  name: Path<T>;
  label: string;
  placeholder?: string;
  type?: "text" | "email" | "password";
  autoComplete?: string;
  disabled?: boolean;
}

export interface AuthSidePanelProps {
  isRegister: boolean;
  onSwitch: () => void;
}

export interface LoginFormProps {
  onSwitch: () => void;
}

export interface RegisterFormProps {
  onSwitch: () => void;
}

export type SignMode = "signIn" | "signUp";

export interface SignFormValues {
  userName: string;
  email: string;
  password: string;
  confirmPassword: string;
}

export interface ApiResponse {
  message?: string;
  error?: string;
}
