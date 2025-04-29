import { Route, Routes } from "react-router-dom";
import { ToastContainer } from "react-toastify";

import { LoginPage } from "./pages/LoginPage";
import { GamePage } from "./pages/GamePage";
import { RegisterPage } from "./pages/RegisterPage";

export const App = () => {
	return (
		<>
			<Routes>
				<Route path="/" element={<LoginPage />} />
				<Route path="/register" element={<RegisterPage />} />
				<Route path="/game" element={<GamePage />} />
			</Routes>
			<ToastContainer />
		</>
	)
}