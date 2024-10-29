# Step 1: Build the React app using a Node.js image
FROM node:14-alpine AS build

# Set working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the app for production
RUN npm run build

# Step 2: Serve the React app using Nginx
FROM nginx:alpine

# Copy the built React app to the Nginx web directory
COPY --from=build /app/build /usr/share/nginx/html

# Expose port 80 to serve the app
EXPOSE 80

# Command to run Nginx
CMD ["nginx", "-g", "daemon off;"]
