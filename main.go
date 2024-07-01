package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const maxColumns = 20

func main() {
    screenWidth := int32(800)
    screenHeight := int32(450)

    rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d camera first person")
    defer rl.CloseWindow()

    camera := rl.Camera{}
    camera.Position = rl.NewVector3(0.0, 2.0, 4.0) // Camera position
    camera.Target = rl.NewVector3(0.0, 2.0, 0.0)   // Camera looking at point
    camera.Up = rl.NewVector3(0.0, 1.0, 0.0)       // Camera up vector (rotation towards target)
    camera.Fovy = 60.0                             // Camera field-of-view Y
    camera.Projection = rl.CameraPerspective             // Camera projection type

    var heights [maxColumns]float32
    var positions [maxColumns]rl.Vector3
    var colors [maxColumns]rl.Color

    for i := 0; i < maxColumns; i++ {
        heights[i] = float32(rl.GetRandomValue(1, 12))
        positions[i] = rl.NewVector3(float32(rl.GetRandomValue(-15, 15)), heights[i]/2.0, float32(rl.GetRandomValue(-15, 15)))
        colors[i] = rl.NewColor(uint8(rl.GetRandomValue(20, 255)), uint8(rl.GetRandomValue(10, 55)), 30, 255)
    }

    rl.DisableCursor() // Limit cursor to relative movement inside the window

    rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

    // Main game loop
    for !rl.WindowShouldClose() {
        rl.UpdateCamera(&camera, rl.CameraFirstPerson) // Update camera

        rl.BeginDrawing()

        rl.ClearBackground(rl.RayWhite)

        rl.BeginMode3D(camera)
		rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(32.0, 32.0), rl.LightGray)
        for i := 0; i < maxColumns; i++ {
            rl.DrawCube(positions[i], 2.0, heights[i], 2.0, colors[i])
            rl.DrawCubeWires(positions[i], 2.0, heights[i], 2.0, rl.Gray)
        }

        rl.EndMode3D()

        rl.EndDrawing()
    }
}