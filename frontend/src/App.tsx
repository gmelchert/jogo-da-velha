import { Route, Routes } from "react-router-dom";
import { ToastContainer } from "react-toastify";

import { LoginPage } from "./pages/LoginPage";
import { GamePage } from "./pages/GamePage";
import { RegisterPage } from "./pages/RegisterPage";
import { HomePage } from "./pages/HomePage";
import { RoomsPage } from "./pages/RoomsPage";

export const App = () => {
	return (
		<>
			<Routes>
				<Route path="/" element={<HomePage />} />
				<Route path="/Login" element={<LoginPage />} />
				<Route path="/register" element={<RegisterPage />} />
				<Route path="/game" element={<GamePage />} />
				<Route path="/rooms" element={<RoomsPage />} />
			</Routes>
			<ToastContainer />
		</>
	)
}