import { ReactNode } from "react";

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  return (
    <>
      <nav className="sticky top-0 z-10 w-full bg-white p-6"></nav>
      <main className="relative min-h-screen bg-komiser-100 pt-10 px-24">{children}</main>
    </>
  );
}

export default Layout;
