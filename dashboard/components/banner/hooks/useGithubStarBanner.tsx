import { useEffect, useState } from 'react';

async function getKomiserGithubStars() {
  try {
    const response = await fetch(
      'https://api.github.com/repos/tailwarden/komiser'
    ).then(res => res.json());
    return response;
  } catch (error) {
    throw new Error(
      'There was an error fetching the GitHub stars from Komiser project.'
    );
  }
}

function useGithubStarBanner() {
  const [displayBanner, setDisplayBanner] = useState(false);
  const [githubStars, setGithubStars] = useState<number>();

  function checkLocalStorageForBannerStatus() {
    if (typeof window !== 'undefined') {
      return localStorage.displayGithubStarBanner;
    }
    return null;
  }

  function dismissBanner() {
    setDisplayBanner(false);
    localStorage.displayGithubStarBanner = 'false';
  }

  useEffect(() => {
    const shouldDisplayBanner = checkLocalStorageForBannerStatus();

    if (shouldDisplayBanner !== 'false') {
      getKomiserGithubStars().then(res => {
        if (!res.stargazers_count) {
          setGithubStars(undefined);
          setDisplayBanner(false);
        } else {
          setGithubStars(res.stargazers_count);
          setDisplayBanner(true);
        }
      });
    }
  }, []);

  return { displayBanner, setDisplayBanner, dismissBanner, githubStars };
}

export default useGithubStarBanner;
