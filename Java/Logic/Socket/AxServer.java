package Java.Logic.Socket;

import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelOption;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;
import io.netty.handler.logging.LogLevel;
import io.netty.handler.logging.LoggingHandler;

public final class AxServer {
    private NioEventLoopGroup masterGroup;
    private NioEventLoopGroup workerGroup;
    private ChannelFuture future;
    static boolean isActive = false;
    private AxServerHandler serverHandler;

    public AxServer(AxServerHandler serverHandler) {
        this.serverHandler = serverHandler;
    }
    public void startServer(int port) throws Exception {
        masterGroup = new NioEventLoopGroup();
        workerGroup = new NioEventLoopGroup();
        try {
            ServerBootstrap b = new ServerBootstrap();
            b.group(masterGroup, workerGroup)
                    .channel(NioServerSocketChannel.class)
                    .option(ChannelOption.SO_BACKLOG, 100)
                    .handler(new LoggingHandler(LogLevel.INFO))
                    .childHandler(new ChannelInitializer<SocketChannel>() {
                        @Override
                        public void initChannel(SocketChannel ch) throws Exception {
                            ChannelPipeline p = ch.pipeline();
                            p.addLast(serverHandler);
                        }
                    })
                    .childOption(ChannelOption.SO_KEEPALIVE, true);
            // Start the server.
            isActive = true;
            future = b.bind(port).sync();
            System.out.println("Starting server on port " + port);
        } catch (Exception e) {
            // Shut down all event loops to terminate all threads.
            e.printStackTrace();
            shutdown();
        }
    }

    public void shutdown() {
        System.out.println("Stopping server");
        try {
            isActive = false;
            if (masterGroup != null) {
                masterGroup.shutdownGracefully().sync();
            }
            if (workerGroup != null) {
                workerGroup.shutdownGracefully().sync();
            }
            if (future != null) {
                future.channel().closeFuture().sync();
            }
            System.out.println("Server stopped");
        }
        catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    public static boolean isActive() {
        return isActive;
    }
}
