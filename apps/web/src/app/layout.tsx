import type { Metadata } from 'next';
import { Plus_Jakarta_Sans } from 'next/font/google';
import './globals.css';
import { StoreProvider } from '@/store/provider';

const plusJakartaSans = Plus_Jakarta_Sans({
  subsets: ['latin'],
  variable: '--font-plus-jakarta',
  display: 'swap',
  weight: ['400', '500', '600', '700', '800'],
});

export const metadata: Metadata = {
  metadataBase: new URL(process.env.NEXT_PUBLIC_APP_URL || 'http://localhost:3000'),
  title: 'Shipfide | Delivering Trust, Every Time',
  description: 'Enterprise parcel delivery, real-time rider tracking, merchant settlements, and hub operations.',
  icons: {
    icon: '/logo.png',
    shortcut: '/logo.png',
    apple: '/logo.png',
  },
  openGraph: {
    title: 'Shipfide | Delivering Trust, Every Time',
    description: 'Enterprise parcel delivery, real-time rider tracking, merchant settlements, and hub operations.',
    images: [{ url: '/logo.png', alt: 'Shipfide Logo' }],
  },
  twitter: {
    card: 'summary_large_image',
    title: 'Shipfide | Delivering Trust, Every Time',
    description: 'Enterprise parcel delivery, real-time rider tracking, merchant settlements, and hub operations.',
    images: ['/logo.png'],
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`h-full antialiased ${plusJakartaSans.variable}`} suppressHydrationWarning>
      <body
        className={`${plusJakartaSans.className} min-h-full flex flex-col bg-[#F9FCFE] text-[#404042] selection:bg-[#DD0033] selection:text-white`}
        suppressHydrationWarning
      >
        <StoreProvider>{children}</StoreProvider>
      </body>
    </html>
  );
}
