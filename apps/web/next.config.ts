import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    images: {
        remotePatterns: [
            {
                protocol: "https",
                hostname: "images.unsplash.com",
            },
        ],
    },
    reactStrictMode: true,
    // Ensure clean CSS & JS asset chunk handling in Next.js App Router
    compiler: {
        removeConsole: process.env.NODE_ENV === "production" ? { exclude: ["error"] } : false,
    },
};

export default nextConfig;
