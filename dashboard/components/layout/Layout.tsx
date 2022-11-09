import { ReactNode } from 'react';
import Navbar from '../navbar/Navbar';

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  return (
    <>
      <Navbar />
      <main className="relative mt-[73px] bg-black-100 py-8 px-8 lg:px-24">
        {children}
      </main>
    </>
  );
}

export default Layout;
