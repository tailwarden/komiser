import classNames from 'classnames';
import Button from '../button/Button';
import type { ToastProps } from './hooks/useToast';

type ToastProp = ToastProps & {
  dismissToast: () => void;
};

function Toast({ hasError, title, message, dismissToast }: ToastProp) {
  return (
    <>
      <div
        className={classNames(
          'fixed bottom-4 left-4 right-4 z-40 flex max-w-2xl animate-fade-in-up items-center justify-between overflow-hidden rounded-lg py-4 px-6 text-black-900 opacity-0 shadow-2xl sm:left-8',
          {
            'bg-error-100': hasError,
            'bg-success-100': !hasError
          }
        )}
      >
        <div
          className={classNames(
            'absolute bottom-0 left-0 h-1 animate-width-to-fit',
            {
              'bg-error-600/60': hasError,
              'bg-success-600/60': !hasError
            }
          )}
        ></div>
        <div className="flex items-center gap-4">
          <div
            className={classNames({
              'text-error-600': hasError,
              'text-success-600': !hasError
            })}
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
                  strokeWidth="1.5"
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
                  strokeWidth="1.5"
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
        <Button style="ghost" onClick={dismissToast}>
          Dismiss
        </Button>
      </div>
    </>
  );
}

export default Toast;
