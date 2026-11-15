import { AuthHeaderProps } from "@/features/auth/types/type";

const AuthHeader = ({ title, description }: AuthHeaderProps) => {
  return (
    <div className="flex flex-col gap-2">
      <h1 className="text-3xl-bold text-primary-text">{title}</h1>
      <p className="text-sm-medium text-secondary-text">{description}</p>
    </div>
  );
};

export default AuthHeader;
