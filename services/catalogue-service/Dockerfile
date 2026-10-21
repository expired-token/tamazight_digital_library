# Stage 1: Dependencies
FROM maven:3.9.9-eclipse-temurin-21-alpine AS dependencies
WORKDIR /app
COPY pom.xml .
RUN mvn dependency:go-offline

# Stage 2: Build
FROM dependencies AS builder
COPY src ./src
RUN mvn clean package -DskipTests

# Stage 3: Security scan
FROM aquasec/trivy:latest AS security
WORKDIR /scan
COPY --from=builder /app/target/*.jar app.jar
RUN trivy fs app.jar

# Stage 4: Runtime
FROM eclipse-temurin:21-jre
WORKDIR /app
COPY --from=builder /app/target/*.jar app.jar

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "app.jar"]