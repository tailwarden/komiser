import { ReactNode } from 'react';
import Banner from '../banner/Banner';
import useGithubStarBanner from '../banner/hooks/useGithubStarBanner';
import Navbar from '../navbar/Navbar';
import LayoutContext from './context/LayoutContext';

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  const { displayBanner, dismissBanner, githubStars } = useGithubStarBanner();

  return (
    <LayoutContext.Provider value={{ displayBanner, dismissBanner }}>
      <Banner githubStars={githubStars} />
      <Navbar />
      <main
        className={`relative ${
          displayBanner ? 'mt-[145px]' : 'mt-[73px]'
        } min-h-screen bg-black-100 p-6 xl:px-8 2xl:px-24`}
      >
        {children}
      </main>
    </LayoutContext.Provider>
  );
}

export default Layout;
