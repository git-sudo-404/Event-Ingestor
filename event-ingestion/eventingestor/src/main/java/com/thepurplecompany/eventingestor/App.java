package com.thepurplecompany.eventingestor;

import com.sun.net.httpserver.HttpServer;
import java.io.IOException;
import java.net.InetSocketAddress;

/** Hello world! */
public class App {
  public static void main(String[] args) throws IOException {
    HttpServer server =
        HttpServer.create(new InetSocketAddress("localhost", 8080), 0);
    server.createContext("/events", new EventsHandler());
    server.start();
    System.out.println("Server started on http://localhost:8080");
  }
}
