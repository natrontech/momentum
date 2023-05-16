import "./globals.css";
import { Inter } from "next/font/google";
import "tw-elements/dist/css/tw-elements.min.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
    title: "Momentum",
    description: "Propel your GitOps workflow with Momentum",
};

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <html lang="en">
            <body className={inter.className}>{children}</body>
        </html>
    );
}
