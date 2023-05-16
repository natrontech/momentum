"use client";

import Button from "./components/general/buttons/Button";

export default function Home() {
    return (
        <div className="App">
            <Button
                text="Click me"
                onClick={() => console.log("Button clicked!")}
            />
        </div>
    );
}
