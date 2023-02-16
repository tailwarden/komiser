import { ReactNode } from 'react';
import Banner from '../banner/Banner';
import useGithubStarBanner from '../banner/hooks/useGithubStarBanner';
import Navbar from '../navbar/Navbar';
import GlobalAppContext from './context/GlobalAppContext';
import useGlobalStats from './hooks/useGlobalStats';

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  const { displayBanner, dismissBanner, githubStars } = useGithubStarBanner();
  const { loading, data, error, hasNoAccounts, fetch } = useGlobalStats();

  return (
    <GlobalAppContext.Provider
      value={{
        displayBanner,
        dismissBanner,
        loading,
        data,
        error,
        hasNoAccounts,
        fetch
      }}
    >
      <Banner githubStars={githubStars} />
      <Navbar />
      <main
        className={`relative ${
          displayBanner
            ? 'mt-[145px] min-h-[calc(100vh-145px)]'
            : 'mt-[73px] min-h-[calc(100vh-73px)]'
        } bg-black-100 p-6 xl:px-8 2xl:px-24`}
      >
        {hasNoAccounts ? <>There is no account</> : children}
      </main>
    </GlobalAppContext.Provider>
  );
}

export default Layout;
