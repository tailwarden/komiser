import { useContext } from 'react';
import formatNumber from '../../utils/formatNumber';
import Button from '../button/Button';
import LayoutContext from '../layout/context/LayoutContext';

type BannerProps = {
  githubStars: number | undefined;
};

function Banner({ githubStars }: BannerProps) {
  const { displayBanner, dismissBanner } = useContext(LayoutContext);

  return (
    <div
      className={`${
        displayBanner ? 'fixed' : 'hidden'
      } top-0 z-10 flex w-full items-center justify-center gap-6 bg-gradient-to-br from-primary to-secondary py-3`}
    >
      <span className="text-sm font-medium text-white">
        Support Komiser by giving us a star on GitHub.
      </span>

      {githubStars && (
        <Button style="bulk-outline">
          Star Komiser {formatNumber(githubStars)}
        </Button>
      )}

      <button
        onClick={dismissBanner}
        className="absolute right-8 cursor-pointer rounded-lg p-3 text-white transition-colors hover:bg-white/10 active:bg-black-900/10"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="none"
          viewBox="0 0 24 24"
        >
          <path
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="1.5"
            d="M8 16l8-8M16 16L8 8"
          ></path>
        </svg>
      </button>
    </div>
  );
}

export default Banner;
