import { ToastProps } from './hooks/useToast';

type ToastProp = ToastProps & {
  dismissToast: () => void;
};

function Toast({ hasError, title, message, dismissToast }: ToastProp) {
  return (
    <>
      <div
        className={`fixed overflow-hidden opacity-0 flex items-center justify-between max-w-2xl z-40 py-4 px-6 bottom-4 left-4 right-4 sm:left-8 rounded-lg shadow-2xl text-black-900 animate-fade-in-up ${
          hasError
            ? `bg-error-100 dark:bg-error-600`
            : `bg-success-100 dark:bg-success-600`
        }`}
      >
        <div
          className={`absolute h-1 bottom-0 left-0 animate-width-to-fit ${
            hasError
              ? `bg-error-600/60 dark:bg-error-700`
              : `bg-success-600/60 dark:bg-success-100/60`
          }`}
        ></div>
        <div className="flex items-center gap-4">
          <div
            className={`${
              hasError
                ? `text-error-600 dark:text-error-700`
                : `text-success-600 dark:text-success-100`
            }`}
          >
            {hasError ? (
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
                  strokeWidth="2"
                  d="M12 22c5.5 0 10-4.5 10-10S17.5 2 12 2 2 6.5 2 12s4.5 10 10 10zM12 8v5M11.995 16h.009"
                ></path>
              </svg>
            ) : (
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
                  strokeWidth="2"
                  d="M20 6L9 17l-5-5"
                ></path>
              </svg>
            )}
          </div>
          <div>
            <p className="text-sm font-medium">{title}</p>
            <p
              className="text-sm text-black-900/60"
              dangerouslySetInnerHTML={{ __html: message }}
            />
          </div>
        </div>
        <div className="w-12"></div>
        <button
          onClick={dismissToast}
          className="bg-transparent hover:bg-black-400/10 active:bg-black-400/20 text-black-900/60 py-4 px-6 text-sm font-medium rounded-lg transition-all"
        >
          Dismiss
        </button>
      </div>
    </>
  );
}

export default Toast;
